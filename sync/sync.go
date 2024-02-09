package sync

import (
	"context"
	"fmt"

	"github.com/leonnicolas/gooki/config"
	"github.com/leonnicolas/gooki/google"
	"github.com/leonnicolas/gooki/nuki"
	"github.com/leonnicolas/gooki/nuki/models"

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
	err = c.SyncUsers(ctx, cfg.UserMatch)
	if err != nil {
		return err
	}

	return nil
}

func (s *syncGSuite) removeDeletedUsers(ctx context.Context) error {
	log.Debug("get deleted users")
	deletedUsers, err := s.google.GetDeletedUsers()
	if err != nil {
		log.Warn("Error Getting Deleted Users", "error", err.Error())
		return err
	}

	for _, u := range deletedUsers {
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

func (s *syncGSuite) SyncNewUsers(ctx context.Context, query string) error {
	log.Debug("get active google users")
	googleUsers, err := s.google.GetUsers(query)
	if err != nil {
		return err
	}

	log.Debug("found google users count=", len(googleUsers))

	for _, u := range googleUsers {
		ll := log.WithField("email", u.PrimaryEmail)
		if s.ignoreUser(u.PrimaryEmail) {
			ll.Debug("ignoring user")
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
		uu, err = s.nukiclient.CreateUser(ctx, &nuki.User{
			Email: &u.PrimaryEmail,
			Name:  ptr(u.Name.DisplayName),
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
	if err := s.removeDeletedUsers(ctx); err != nil {
		log.WithError(err).Error("failed to remove deleted users")
		return err
	}

	if err := s.SyncNewUsers(ctx, query); err != nil {
		log.WithError(err).Error("failed to sync new users")
		return err
	}

	return nil
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

func (s *syncGSuite) ignoreUser(name string) bool {
	for _, u := range s.cfg.IgnoreUsers {
		if u == name {
			return true
		}
	}

	return false
}
