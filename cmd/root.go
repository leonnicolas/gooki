// Copyright (c) 2020, Amazon.com, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cmd ...
package cmd

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/metalmatze/signal/internalserver"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/leonnicolas/gooki/config"
	"github.com/leonnicolas/gooki/sync"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Version: "dev",
	Use:     "gooki",
	Short:   "sync nuki users from google groups",
	Long: `A command line tool to enable you to synchronise your Google
Apps (Google Workspace) users to nuki users`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		log.Info("period ", cfg.SyncPeriod)
		if cfg.SyncPeriod > 0 {
			reg := prometheus.NewRegistry()
			cv := prometheus.NewCounterVec(prometheus.CounterOpts{
				Name: "gooki_syncs_total",
				Help: "Total number of syncs",
			}, []string{"status"})
			reg.MustRegister(
				collectors.NewGoCollector(),
				collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
				cv,
			)
			g := &run.Group{}
			g.Add(run.SignalHandler(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM))
			{
				h := internalserver.NewHandler(
					internalserver.WithName("Internal - advo-assist"),
					internalserver.WithPrometheusRegistry(reg),
					internalserver.WithPProf(),
				)
				l, err := net.Listen("tcp", cfg.ListenInternal)
				if err != nil {
					return fmt.Errorf("failed to listen on %s: %w", cfg.ListenInternal, err)
				}

				g.Add(func() error {
					if err := http.Serve(l, h); err != nil && err != http.ErrServerClosed {
						return fmt.Errorf("error: internal server exited unexpectedly: %v", err)
					}
					return nil
				}, func(error) {
					l.Close()
				})
			}
			g.Add(func() error {
				for {
					err := sync.DoSync(ctx, cfg)
					if err != nil {
						log.WithError(err).Error("failed to sync")
						cv.WithLabelValues("error").Inc()
					} else {
						cv.WithLabelValues("success").Inc()
					}
					select {
					case <-time.After(cfg.SyncPeriod):
					case <-ctx.Done():
						return context.Canceled
					}
				}
			}, func(error) {
				cancel()
			})
			return g.Run()

		}
		err := sync.DoSync(ctx, cfg)
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute is the entry point of the command. If we are
// running inside of AWS Lambda, we use the Lambda
// execution path.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// init config
	cfg = config.New()

	// initialize cobra
	cobra.OnInitialize(initConfig)
	addFlags(rootCmd, cfg)

	rootCmd.SetVersionTemplate(fmt.Sprintf("%s, commit %s, built at %s by %s\n", version, commit, date, builtBy))

	// silence on the root cmd
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// allow to read in from environment
	viper.SetEnvPrefix("gooki")
	viper.AutomaticEnv()

	appEnvVars := []string{
		"google_admin",
		"google_credentials",
		"log_level",
		"log_format",
		"ignore_users",
		"user_match",
	}

	for _, e := range appEnvVars {
		if err := viper.BindEnv(e); err != nil {
			log.Fatalf(errors.Wrap(err, "cannot bind environment variable").Error())
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(errors.Wrap(err, "cannot unmarshal config").Error())
	}

	f, err := os.Open(cfg.GoogleCredentials)
	if err != nil {
		log.Fatal("faile to open google credentials file", err.Error())
	}
	buf, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("faile to read credentials file", err.Error())
	}

	cfg.GoogleCredentials = string(buf)

	logConfig(cfg)
}

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	rootCmd.PersistentFlags().StringVarP(&cfg.GoogleCredentials, "google-admin", "a", config.DefaultGoogleCredentials, "path to find credentials file for Google Workspace")
	rootCmd.PersistentFlags().BoolVarP(&cfg.Debug, "debug", "d", config.DefaultDebug, "enable verbose / debug logging")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogFormat, "log-format", "", config.DefaultLogFormat, "log format")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogLevel, "log-level", "", config.DefaultLogLevel, "log level")
	rootCmd.Flags().StringVarP(&cfg.GoogleCredentials, "google-credentials", "c", config.DefaultGoogleCredentials, "path to Google Workspace credentials file")
	rootCmd.Flags().StringVarP(&cfg.GoogleAdmin, "google-admin", "u", "", "Google Workspace admin user email")
	rootCmd.Flags().StringSliceVar(&cfg.IgnoreUsers, "ignore-users", []string{}, "ignores these Google Workspace users")
	rootCmd.Flags().StringVarP(&cfg.Query, "query", "q", "", "Google Workspace filter query parameter, example: 'name:John* email:admin*', see: https://developers.google.com/admin-sdk/directory/v1/guides/search-users and https://developers.google.com/admin-sdk/directory/v1/guides/search-groups")
	rootCmd.Flags().StringSliceVar(&cfg.IgnoreGroups, "ignore-groups", []string{}, "ignores these Google Workspace groups")
	rootCmd.Flags().StringVarP(&cfg.SyncMethod, "sync-method", "s", config.DefaultSyncMethod, "Sync method to use (users_groups|groups)")
	rootCmd.Flags().StringVarP(&cfg.NukiBearerToken, "nuki-token", "t", "", "Nuki Bearer Token for authentication with the nuki API.")
	rootCmd.Flags().Int64Var(&cfg.SmartlockID, "smartlock-id", 0, "Nuki Smartlock ID; if 0 no smartlock auths will be created.")
	rootCmd.Flags().DurationVarP(&cfg.SyncPeriod, "sync-period", "p", 0, "Time between two syncronizations, will only run once if set to 0")
	rootCmd.Flags().StringVar(&cfg.ListenInternal, "listen", ":8080", "Address to listen on for internal server")
}

func logConfig(cfg *config.Config) {
	// reset log format
	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if cfg.Debug {
		cfg.LogLevel = "debug"
	}

	// set the configured log level
	if level, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(level)
	}
}
