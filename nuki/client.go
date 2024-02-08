package nuki

import (
	"context"
	"errors"
	"fmt"
	"sync"

	httptransport "github.com/go-openapi/runtime/client"
	log "github.com/sirupsen/logrus"

	nukiclient "github.com/leonnicolas/gooki/nuki/client"
	"github.com/leonnicolas/gooki/nuki/client/account_user"
	"github.com/leonnicolas/gooki/nuki/client/smartlock_auth"
	"github.com/leonnicolas/gooki/nuki/models"
)

var ErrUserNotFound = errors.New("user not found")

type Client struct {
	nukiClient *nukiclient.NukiAPI
	beareToken string
}

type User models.AccountUser

func (u *User) GetEmail() *string {
	return u.Email
}

func (u *User) GetName() *string {
	return u.Name
}

func (u *User) GetID() *int32 {
	return u.AccountUserID
}

func New(nukiAPI *nukiclient.NukiAPI, beareToken string) *Client {
	return &Client{
		nukiClient: nukiAPI,
		beareToken: beareToken,
	}
}

func NewDefault(bearerToken string) *Client {
	return New(
		nukiclient.NewHTTPClientWithConfig(nil, &nukiclient.TransportConfig{
			Host:    "api.nuki.io",
			Schemes: []string{"https"},
		}),
		bearerToken,
	)
}

func NewCached(nukiAPI *nukiclient.NukiAPI, beareToken string) *CachedClient {
	return &CachedClient{
		client:     New(nukiAPI, beareToken),
		nukiClient: nukiAPI,
		beareToken: beareToken,
	}
}

func NewCachedDefault(bearerToken string) *CachedClient {
	return NewCached(
		nukiclient.NewHTTPClientWithConfig(nil, &nukiclient.TransportConfig{
			Host:    "api.nuki.io",
			Schemes: []string{"https"},
		}),
		bearerToken,
	)
}

type CachedClient struct {
	client     *Client
	nukiClient *nukiclient.NukiAPI
	beareToken string

	mux                sync.Mutex
	userCache          map[string]*User
	mux0               sync.Mutex
	smartlockAuthCache map[int32]*models.SmartlockAuth
}

func (c *CachedClient) CreateSmartlockAuth(ctx context.Context, sa *models.SmartlocksAuthCreate) error {
	if err := c.client.CreateSmartlockAuth(ctx, sa); err != nil {
		return err
	}
	c.mux0.Lock()
	defer c.mux0.Unlock()
	c.smartlockAuthCache[sa.AccountUserID] = &models.SmartlockAuth{
		AccountUserID:   sa.AccountUserID,
		AllowedWeekDays: sa.AllowedWeekDays,
		Name:            sa.Name,
		Type:            sa.Type,
		RemoteAllowed:   sa.RemoteAllowed,
	}
	return nil
}

func (c *CachedClient) CreateUser(ctx context.Context, u *User) (*User, error) {
	r, err := c.client.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	c.userCache[*r.Email] = r
	return r, nil
}

func (c *CachedClient) DeleteUser(ctx context.Context, u *User) error {
	err := c.client.DeleteUser(ctx, u)
	if err != nil {
		return err
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	delete(c.userCache, *u.Email)
	return nil
}

// CreateSmartlockAuth is async
func (c *Client) CreateSmartlockAuth(ctx context.Context, sa *models.SmartlocksAuthCreate) error {
	log.WithField("name", *sa.Name).Debug("creating smart lock auth")
	_, err := c.nukiClient.SmartlockAuth.SmartlocksAuthsResourcePutPut(
		&smartlock_auth.SmartlocksAuthsResourcePutPutParams{
			Context: ctx,
			Body:    sa,
		},
		httptransport.BearerToken(c.beareToken),
	)
	return err
}

func (c *CachedClient) FindSmartlockAuth(ctx context.Context, smartlockID int64, accoutnUserId int32) (*models.SmartlockAuth, error) {
	log := log.WithField("accoutnUserId", accoutnUserId).WithField("smartlockID", smartlockID)
	if c.smartlockAuthCache != nil {
		if u, ok := c.smartlockAuthCache[accoutnUserId]; ok {
			return u, nil
		}
		return nil, ErrUserNotFound
	}
	c.mux0.Lock()
	defer c.mux0.Unlock()

	c.smartlockAuthCache = make(map[int32]*models.SmartlockAuth)

	userFindRes, err := c.nukiClient.SmartlockAuth.SmartlockAuthsResourceGetGet(
		&smartlock_auth.SmartlockAuthsResourceGetGetParams{
			SmartlockID: smartlockID,
			Types:       ptr("0"),
			Context:     ctx,
		},
		httptransport.BearerToken(c.beareToken),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get smart lock auths: %w", err)
	}
	log.WithField("count", len(userFindRes.Payload)).Debug("found smart lock auths")

	for _, u := range userFindRes.Payload {
		c.smartlockAuthCache[u.AccountUserID] = u
	}
	return c.FindSmartlockAuth(ctx, smartlockID, accoutnUserId)
}

func (c *CachedClient) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	log := log.WithField("operation", "find user").WithField("cache", true).WithField("email", email)

	if c.userCache != nil {
		if u, ok := c.userCache[email]; ok {
			return u, nil
		}
		return nil, ErrUserNotFound
	}
	c.mux.Lock()
	defer c.mux.Unlock()

	c.userCache = make(map[string]*User)

	offset := int64(0)
	limit := int64(10)

	for {
		log.WithField("offset", offset).WithField("limit", limit).Debug("fetching users")
		userFindRes, err := c.nukiClient.AccountUser.AccountUsersResourceGetGet(
			&account_user.AccountUsersResourceGetGetParams{
				Context: ctx,
				Limit:   ptr(limit),
				Offset:  ptr(offset),
			},
			httptransport.BearerToken(c.beareToken),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get user %w", err)
		}
		log.Debug("found users", "count", len(userFindRes.Payload))
		if len(userFindRes.Payload) == 0 {
			log.Debug("no more users found")
			break
		}

		for _, u := range userFindRes.Payload {
			c.userCache[*u.Email] = (*User)(u)
		}
		offset += int64(len(userFindRes.Payload))
	}
	return c.FindUserByEmail(ctx, email)
}

func ptr[E any](e E) *E {
	return &e
}

func (c *Client) CreateUser(ctx context.Context, u *User) (*User, error) {
	log.Info("CreateUser")
	userCreateRes, err := c.nukiClient.AccountUser.AccountUsersResourcePutPut(
		&account_user.AccountUsersResourcePutPutParams{
			Context: ctx,
			Body: &models.AccountUserCreate{
				Email:    u.GetEmail(),
				Language: "de",
				Name:     u.GetName(),
			},
		},
		httptransport.BearerToken(c.beareToken),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user %w", err)
	}

	return (*User)(userCreateRes.Payload), nil
}

func (c *Client) DeleteUser(ctx context.Context, u *User) error {
	log.Info("DeleteUser")
	_, err := c.nukiClient.AccountUser.AccountUserResourceDeleteDelete(
		&account_user.AccountUserResourceDeleteDeleteParams{
			Context:       ctx,
			AccountUserID: int64(*u.GetID()),
		},
		httptransport.BearerToken(c.beareToken),
	)
	if err != nil {
		return fmt.Errorf("failed to delete user %w", err)
	}

	return nil
}

func (c *Client) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	log.Info("Find user", "email", email)
	userFindRes, err := c.nukiClient.AccountUser.AccountUsersResourceGetGet(
		&account_user.AccountUsersResourceGetGetParams{
			Context: ctx,
			Email:   &email,
		},
		httptransport.BearerToken(c.beareToken),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user %w", err)
	}
	if len(userFindRes.Payload) > 1 {
		return nil, fmt.Errorf("found more than one user with email %s", email)
	}
	if len(userFindRes.Payload) == 0 {
		return nil, ErrUserNotFound
	}

	return (*User)(userFindRes.Payload[0]), nil
}
