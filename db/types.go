package db

type Dialect string

const (
	DialectMariaDB  Dialect = "mariadb"
	DialectMySQL    Dialect = "mysql"
	DialectPostgres Dialect = "postgres"
	DialectSQLite3  Dialect = "sqlite3"
)
