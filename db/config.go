package db

import (
	"runtime"
	"time"
)

var ConfigDefault = Config{
	Dialect:  DialectMySQL,
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "",
	Database: "db",
	Timezone: "UTC",
	Debug:    false,

	ConnMaxIdleTime: 3 * time.Minute,
	ConnMaxLifetime: 30 * time.Minute,
	MaxOpenConns:    runtime.NumCPU() * 4,
	MaxIdleConns:    runtime.NumCPU() * 2,
}

type Config struct {
	Dialect  Dialect
	DSN      string
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Timezone string
	Debug    bool

	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

// Helper function to set default values
func configDefault(config Config) Config {
	// Override default config
	if config.Dialect == "" {
		config.Dialect = ConfigDefault.Dialect
	}
	if config.Host == "" {
		config.Host = ConfigDefault.Host
	}
	if config.Port == 0 {
		config.Port = ConfigDefault.Port
	}
	if config.User == "" {
		config.User = ConfigDefault.User
	}
	if config.Password == "" {
		config.Password = ConfigDefault.Password
	}
	if config.Database == "" {
		config.Database = ConfigDefault.Database
	}

	if config.ConnMaxIdleTime == 0 {
		config.ConnMaxIdleTime = ConfigDefault.ConnMaxIdleTime
	}
	if config.ConnMaxLifetime == 0 {
		config.ConnMaxLifetime = ConfigDefault.ConnMaxLifetime
	}
	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = ConfigDefault.MaxOpenConns
	}
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = ConfigDefault.MaxIdleConns
	}

	return config
}
