package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	mysql "github.com/go-sql-driver/mysql"
)

func NewSQL(config Config) (*sql.DB, error) {
	config = configDefault(config)
	dsn := makeDSN(config)

	mysql.SetLogger(log.Default())
	db, err := sql.Open(string(config.Dialect), dsn)
	if err != nil {
		return nil, err
	}

	switch config.Dialect {
	case DialectMySQL, DialectPostgres:
		db.SetConnMaxIdleTime(3 * time.Minute)
		db.SetMaxOpenConns(16)
		db.SetMaxIdleConns(16)
	case DialectSQLite3:
		db.SetMaxOpenConns(1)
	}

	return db, nil
}

func makeDSN(cfg Config) string {
	if cfg.DSN != "" {
		return cfg.DSN
	}

	switch cfg.Dialect {
	case DialectMySQL:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=true&loc=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, url.QueryEscape(cfg.Timezone))
	case DialectPostgres:
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=prefer password=%s TimeZone=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.Timezone)
	case DialectSQLite3:
		return fmt.Sprintf("file:%s?cache=shared&mode=rwc&_loc=%s", cfg.Database, url.QueryEscape(cfg.Timezone))
	default:
		return ""
	}
}
