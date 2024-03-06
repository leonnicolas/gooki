package sync

import (
	"context"
	"fmt"

	"github.com/leonnicolas/gooki/config"
	"github.com/leonnicolas/gooki/google"
	"github.com/leonnicolas/gooki/nuki"
	"github.com/leonnicolas/gooki/nuki/models"
	admin "google.golang.org/api/admin/directory/v1"

	log "github.com/sirupsen/logrus"
)

type Client interface {
	CreateUser(context.Context, *nuki.User) (*nuki.User, error)
	FindUserByEmail(context.Context, string) (*nuki.User, error)
	FindSmartlockAuth(ctx context.Context, smartlockID int64, accountUserID int32) (*models.SmartlockAuth, error)
	CreateSmartlockAuth(context.Context, *models.SmartlocksAuthCreate) error
	DeleteUser(context.Context, *nuki.User) error
}

type User interface {
	GetEmail() *string
	GetName() *string
	GetID() *int32
}

type SyncGSuite interface {
	SyncUsers(context.Context, string) error
}

type syncGSuite struct {
	nukiclient Client
	google     google.Client
	cfg        *config.Config
}

func New(cfg *config.Config, a Client, g google.Client) SyncGSuite {
	return &syncGSuite{
		nukiclient: a,
		google:     g,
		cfg:        cfg,
	}
}

var _ User = &nuki.User{}

func DoSync(ctx context.Context, cfg *config.Config) error {
	log.Info("Syncing google users to nuki")

	creds := []byte(cfg.GoogleCredentials)

	googleClient, err := google.NewClient(ctx, cfg.GoogleAdmin, creds)
	if err != nil {
		return fmt.Errorf("failed to initialize google client %w", err)
	}

	nukiClient := nuki.NewCachedDefault(cfg.NukiBearerToken)

	c := New(cfg, nukiClient, googleClient)

	log.Info("syncing")
	err = c.SyncUsers(ctx, cfg.Query)
	if err != nil {
		return err
	}

	return nil
}

func (s *syncGSuite) deletedUsers(ctx context.Context) ([]*admin.User, error) {
	log.Debug("get deleted users")
	deletedUsers, err := s.google.GetDeletedUsers()
	if err != nil {
		log.Warn("Error Getting Deleted Users", "error", err.Error())
		return nil, err
	}
	return deletedUsers, nil
}

func (s *syncGSuite) googleUsers(ctx context.Context, query string) ([]*admin.User, error) {
	var googleUsers []*admin.User
	log.Debug("get active google users")
	if s.cfg.SyncMethod == config.DefaultSyncMethod {
		log.WithField("query", query).Info("get google groups")
		googleGroups, err := s.google.GetGroups(query)
		if err != nil {
			return nil, err
		}
		filteredGoogleGroups := []*admin.Group{}
		for _, g := range googleGroups {
			if s.ignoreGroup(g.Email) {
				log.WithField("group", g.Email).Debug("ignoring group")
				continue
			}
			filteredGoogleGroups = append(filteredGoogleGroups, g)
		}
		googleUsers, _, err = s.getGoogleGroupsAndUsers(filteredGoogleGroups)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		googleUsers, err = s.google.GetUsers(query)
		if err != nil {
			return nil, err
		}
	}

	log.Debug("found google users count=", len(googleUsers))
	return googleUsers, nil
}

func (s *syncGSuite) createNukiUsers(ctx context.Context, googleUsers []*admin.User) error {
	for _, u := range googleUsers {
		ll := log.WithField("email", u.PrimaryEmail)
		if s.ignoreUser(u.PrimaryEmail) {
			ll.Debug("ignoring user")
			continue
		}

		if u.Suspended {
			//	deletedUsers = append(deletedUsers, u)
			continue
		}

		ll.Debug("finding user")
		uu, err := s.nukiclient.FindUserByEmail(ctx, u.PrimaryEmail)
		if err != nil && err != nuki.ErrUserNotFound {
			return fmt.Errorf("failed to find user: %w", err)
		}
		if uu != nil {
			ll.Debug("user exists")
			if s.cfg.SmartlockID == 0 {
				continue
			}
			if err := s.createAuthUser(ctx, uu); err != nil {
				ll.WithError(err).Error("failed to create auth user")
			}
			continue
		}

		ll.Info("creating user")
		name := u.Name.DisplayName
		if name == "" {
			name = u.Name.FullName
		}
		if name == "" {
			log.Warn("skipping user without name")
			continue
		}
		uu, err = s.nukiclient.CreateUser(ctx, &nuki.User{
			Email: &u.PrimaryEmail,
			Name:  ptr(name),
		})
		if err != nil {
			return err
		}
		if s.cfg.SmartlockID == 0 {
			continue
		}
		ll.Info("creating auth user")
		if err := s.createAuthUser(ctx, uu); err != nil {
			ll.WithError(err).Error("failed to create auth user")
		}
	}
	return nil
}

func (s *syncGSuite) SyncUsers(ctx context.Context, query string) error {
	deletedUsers, err := s.deletedUsers(ctx)
	if err != nil {
		return err
	}
	googleUsers, err := s.googleUsers(ctx, query)
	if err != nil {
		return err
	}

	if err := s.createNukiUsers(ctx, googleUsers); err != nil {
		log.WithError(err).Error("failed to sync new users")
		return err
	}

	for _, u := range googleUsers {
		if u.Suspended {
			deletedUsers = append(deletedUsers, u)
		}
	}

	if err := s.deleteUsers(ctx, deletedUsers); err != nil {
		return fmt.Errorf("failed to delete nuki users: %w", err)
	}

	return nil
}

func (s *syncGSuite) deleteUsers(ctx context.Context, googleUsers []*admin.User) error {
	for _, u := range googleUsers {
		log := log.WithField("email", u.PrimaryEmail)
		log.Info("deleting google user")

		uu, err := s.nukiclient.FindUserByEmail(ctx, u.PrimaryEmail)
		if err != nuki.ErrUserNotFound && err != nil {
			log.Warn("Error deleting google user")
			return err
		}

		if err == nuki.ErrUserNotFound {
			log.Debug("User already deleted")
			continue
		}
		err = s.nukiclient.DeleteUser(ctx, uu)
		if err != nil {
			log.Warn("Error deleting user")
			return err
		}
	}
	return nil
}

// getGoogleGroupsAndUsers return a list of google users members of googleGroups
// and a map of google groups and its users' list
func (s *syncGSuite) getGoogleGroupsAndUsers(googleGroups []*admin.Group) ([]*admin.User, map[string][]*admin.User, error) {
	gUsers := make([]*admin.User, 0)
	gGroupsUsers := make(map[string][]*admin.User)

	gUniqUsers := make(map[string]*admin.User)

	for _, g := range googleGroups {

		log := log.WithFields(log.Fields{"group": g.Name})

		if s.ignoreGroup(g.Email) {
			log.Debug("ignoring group")
			continue
		}

		log.Debug("get group members from google")
		groupMembers, err := s.google.GetGroupMembers(g)
		if err != nil {
			return nil, nil, err
		}

		log.Debug("get users")
		membersUsers := make([]*admin.User, 0)

		for _, m := range groupMembers {

			if s.ignoreUser(m.Email) {
				log.WithField("id", m.Email).Debug("ignoring user")
				continue
			}

			log.WithField("id", m.Email).Debug("get user")
			q := fmt.Sprintf("email:%s", m.Email)
			u, err := s.google.GetUsers(q) // TODO: implement GetUser(m.Email)
			if err != nil {
				return nil, nil, err
			}
			if len(u) == 0 {
				log.WithField("id", m.Email).Warn("missing user")
				continue
			}

			membersUsers = append(membersUsers, u[0])

			_, ok := gUniqUsers[m.Email]
			if !ok {
				gUniqUsers[m.Email] = u[0]
			}
		}
		gGroupsUsers[g.Name] = membersUsers
	}

	for _, user := range gUniqUsers {
		gUsers = append(gUsers, user)
	}

	return gUsers, gGroupsUsers, nil
}

func (s *syncGSuite) createAuthUser(ctx context.Context, u *nuki.User) error {
	_, err := s.nukiclient.FindSmartlockAuth(ctx, s.cfg.SmartlockID, *u.AccountUserID)
	if err != nil && err != nuki.ErrUserNotFound {
		return err
	}
	if err == nuki.ErrUserNotFound {
		err := s.nukiclient.CreateSmartlockAuth(ctx, &models.SmartlocksAuthCreate{
			AccountUserID: *u.AccountUserID,
			SmartlockIds:  []int64{s.cfg.SmartlockID},
			RemoteAllowed: ptr(true),
			Type:          0,
			Name:          u.Name,
		})
		if err != nil {
			return fmt.Errorf("failed to create smartlock auth %w", err)
		}
	}
	return nil
}

func ptr[E any](e E) *E {
	return &e
}

func (s *syncGSuite) ignoreGroup(name string) bool {
	for _, g := range s.cfg.IgnoreGroups {
		if g == name {
			return true
		}
	}

	return false
}

func (s *syncGSuite) ignoreUser(name string) bool {
	for _, u := range s.cfg.IgnoreUsers {
		if u == name {
			return true
		}
	}

	return false
}
