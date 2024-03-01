package sync

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/leonnicolas/gooki/config"
	"github.com/leonnicolas/gooki/google"
	"github.com/leonnicolas/gooki/nuki"
	"github.com/stretchr/testify/mock"
	admin "google.golang.org/api/admin/directory/v1"
)

func TestCreateUsers(t *testing.T) {
	type testConfig struct {
		cfg    *config.Config
		nuki   *MockClient
		google *google.MockClient
	}

	query := "q"
	for _, scenario := range []struct {
		name         string
		tc           *testConfig
		expectations func(*testConfig)
		expectedErr  string
	}{
		{
			name: "no users",
			expectations: func(tc *testConfig) {
				tc.google.On("GetUsers", query).Return([]*admin.User{}, nil)
			},
		},
		{
			name: "one user, nuki api error",
			expectations: func(tc *testConfig) {
				tc.google.On("GetUsers", query).Return([]*admin.User{{PrimaryEmail: "email@host.domain", Name: &admin.UserName{DisplayName: "DisplayName"}}}, nil)

				tc.nuki.
					On("FindUserByEmail", mock.Anything, "email@host.domain").
					Return(nil, fmt.Errorf("nuki api error"))
			},
			expectedErr: "nuki api error",
		},
		{
			name: "one user, nuki user found",
			expectations: func(tc *testConfig) {
				tc.google.On("GetUsers", query).Return([]*admin.User{{PrimaryEmail: "email@host.domain", Name: &admin.UserName{DisplayName: "DisplayName"}}}, nil)

				user := &nuki.User{
					AccountID: ptr(int32(12345)),
				}
				tc.nuki.
					On("FindUserByEmail", mock.Anything, "email@host.domain").
					Return(user, fmt.Errorf("nuki api error"))
			},
			expectedErr: "nuki api error",
		},
		{
			name: "one user, nuki user not found, creation error",
			expectations: func(tc *testConfig) {
				gu := &admin.User{PrimaryEmail: "email@host.domain", Name: &admin.UserName{DisplayName: "DisplayName"}}
				tc.google.On("GetUsers", query).Return([]*admin.User{gu}, nil)

				tc.nuki.
					On("FindUserByEmail", mock.Anything, gu.PrimaryEmail).
					Return(nil, nuki.ErrUserNotFound)
				tc.nuki.On("CreateUser", mock.Anything, &nuki.User{
					Email: &gu.PrimaryEmail,
					Name:  &gu.Name.DisplayName,
				}).Return(nil, fmt.Errorf("nuki api error"))
			},
			expectedErr: "nuki api error",
		},
		{
			name: "one user, nuki user not found",
			expectations: func(tc *testConfig) {
				gu := &admin.User{PrimaryEmail: "email@host.domain", Name: &admin.UserName{DisplayName: "DisplayName"}}
				tc.google.On("GetUsers", query).Return([]*admin.User{gu}, nil)

				tc.nuki.
					On("FindUserByEmail", mock.Anything, gu.PrimaryEmail).
					Return(nil, nuki.ErrUserNotFound)
				tc.nuki.On("CreateUser", mock.Anything, &nuki.User{
					Email: &gu.PrimaryEmail,
					Name:  &gu.Name.DisplayName,
				}).Return(&nuki.User{}, nil)
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			if scenario.tc == nil {
				c := config.New()
				c.SyncMethod = "users"
				scenario.tc = &testConfig{
					c,
					&MockClient{},
					&google.MockClient{},
				}
			}
			scenario.expectations(scenario.tc)
			scenario.tc.google.On("GetDeletedUsers").Return([]*admin.User{}, nil)

			sync := &syncGSuite{scenario.tc.nuki, scenario.tc.google, scenario.tc.cfg}

			ctx := context.Background()
			err := sync.SyncUsers(ctx, query)
			if err != nil {
				if scenario.expectedErr == "" {
					t.Errorf("unexpected error removing deleted users: %v", err)
				}
				if !strings.Contains(err.Error(), scenario.expectedErr) {
					t.Errorf("unexpected error removing deleted users: %v; got: %v", scenario.expectedErr, err)
				}
			} else if scenario.expectedErr != "" {
				t.Errorf("expected error removing deleted users: %v", scenario.expectedErr)
			}
		})
	}
}

func TestRemoveDeletedUsers(t *testing.T) {
	type testConfig struct {
		cfg    *config.Config
		nuki   *MockClient
		google *google.MockClient
	}

	for _, scenario := range []struct {
		name         string
		tc           *testConfig
		expectations func(*testConfig)
		expectedErr  string
	}{
		{
			name: "no deleted users",
			expectations: func(tc *testConfig) {
				tc.google.On("GetDeletedUsers").Return([]*admin.User{}, nil)
			},
		},
		{
			name: "one deleted user, that doesn't exist in nuki",
			expectations: func(tc *testConfig) {
				tc.google.
					On("GetDeletedUsers").
					Return([]*admin.User{{PrimaryEmail: "email@host.domain"}}, nil)
				tc.nuki.
					On("FindUserByEmail", mock.Anything, "email@host.domain").
					Return(nil, nuki.ErrUserNotFound)
			},
		},
		{
			name: "one deleted user, nuki api error",
			expectations: func(tc *testConfig) {
				tc.google.
					On("GetDeletedUsers").
					Return([]*admin.User{{PrimaryEmail: "email@host.domain"}}, nil)

				tc.nuki.
					On("FindUserByEmail", mock.Anything, "email@host.domain").
					Return(nil, fmt.Errorf("nuki api error"))
			},
			expectedErr: "nuki api error",
		},
		{
			name: "one deleted user, nuki api error on delete",
			expectations: func(tc *testConfig) {
				tc.google.
					On("GetDeletedUsers").
					Return([]*admin.User{{PrimaryEmail: "email@host.domain"}}, nil)

				user := &nuki.User{
					AccountID: ptr(int32(12345)),
				}
				tc.nuki.
					On("FindUserByEmail", mock.Anything, "email@host.domain").
					Return(user, nil)
				tc.nuki.
					On("DeleteUser", mock.Anything, user).
					Return(fmt.Errorf("nuki api error"))
			},
			expectedErr: "nuki api error",
		},
		{
			name: "one deleted user",
			expectations: func(tc *testConfig) {
				tc.google.
					On("GetDeletedUsers").
					Return([]*admin.User{{PrimaryEmail: "email@host.domain"}}, nil)

				user := &nuki.User{
					AccountID: ptr(int32(12345)),
				}
				tc.nuki.
					On("FindUserByEmail", mock.Anything, "email@host.domain").
					Return(user, nil)
				tc.nuki.
					On("DeleteUser", mock.Anything, user).
					Return(nil)
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			if scenario.tc == nil {
				c := config.New()
				c.SyncMethod = "users"
				scenario.tc = &testConfig{
					c,
					&MockClient{},
					&google.MockClient{},
				}
			}
			scenario.expectations(scenario.tc)
			scenario.tc.google.On("GetUsers", "").Return([]*admin.User{}, nil)

			sync := &syncGSuite{scenario.tc.nuki, scenario.tc.google, scenario.tc.cfg}

			ctx := context.Background()
			err := sync.SyncUsers(ctx, "")
			if err != nil {
				if scenario.expectedErr == "" {
					t.Errorf("unexpected error removing deleted users: %v", err)
				}
				if !strings.Contains(err.Error(), scenario.expectedErr) {
					t.Errorf("unexpected error removing deleted users: %v; got: %v", scenario.expectedErr, err)
				}
			} else if scenario.expectedErr != "" {
				t.Errorf("expected error removing deleted users: %v", scenario.expectedErr)
			}
		})
	}
}
