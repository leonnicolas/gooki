// Package config ...
package config

import "time"

// Config ...
type Config struct {
	// Verbose toggles the verbosity
	Debug bool
	// LogLevel is the level with with to log for this config
	LogLevel string `mapstructure:"log_level"`
	// LogFormat is the format that is used for logging
	LogFormat string `mapstructure:"log_format"`
	// GoogleCredentials ...
	GoogleCredentials string `mapstructure:"google_credentials"`
	// GoogleAdmin ...
	GoogleAdmin string `mapstructure:"google_admin"`
	// UserMatch ...
	UserMatch string `mapstructure:"user_match"`
	// Ignore users ...
	IgnoreUsers     []string `mapstructure:"ignore_users"`
	NukiBearerToken string   `mapstructure:"nuki_beare_token"`
	// SyncPeriod the timeout between two syncronizations
	// Will only run once if set to 0
	SyncPeriod     time.Duration `mapstructure:"sync_period"`
	ListenInternal string        `mapstructure:"listen_internal"`
	SmartlockID    int64         `mapstructure:"smartlock_id"`
}

const (
	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = "info"
	// DefaultLogFormat is the default format of the logger
	DefaultLogFormat = "text"
	// DefaultDebug is the default debug status.
	DefaultDebug = false
	// DefaultGoogleCredentials is the default credentials path
	DefaultGoogleCredentials = "credentials.json"
	// DefaultSyncMethod is the default sync method to use.
	DefaultSyncMethod = "groups"
)

// New returns a new Config
func New() *Config {
	return &Config{
		Debug:             DefaultDebug,
		LogLevel:          DefaultLogLevel,
		LogFormat:         DefaultLogFormat,
		GoogleCredentials: DefaultGoogleCredentials,
	}
}
