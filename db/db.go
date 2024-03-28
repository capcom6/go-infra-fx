package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func New(params Params) (*gorm.DB, error) {
	cfgGorm := makeConfig(params)

	switch params.Config.Dialect {
	case DialectMySQL:
		return gorm.Open(mysql.New(mysql.Config{Conn: params.SQL}), cfgGorm)
	case DialectSQLite3:
		return gorm.Open(sqlite.Open(params.Config.Database), cfgGorm)
	case DialectPostgres:
		return gorm.Open(postgres.New(postgres.Config{Conn: params.SQL}), cfgGorm)
	}

	return nil, fmt.Errorf("unsupported dialect: %s", params.Config.Dialect)
}

func makeConfig(params Params) *gorm.Config {
	log := zapgorm2.New(params.Logger)
	log.LogLevel = logger.Info
	if params.Config.Debug {
		log.LogLevel = logger.Info + 1
	}
	log.SetAsDefault()

	return &gorm.Config{
		Logger: log,
	}
}
