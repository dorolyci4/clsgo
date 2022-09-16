package database

import (
	"gorm.io/gorm"
	// GORM drivers
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
)

// "test.db"
// "file::memory:?cache=shared"
func Sqlite(dsn string) gorm.Dialector {
	return sqlite.Open(dsn)
}

// dsn := "username:password@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
func MySQL(dsn string) gorm.Dialector {
	return mysql.Open(dsn)
}

// dsn := "sqlserver://username:password@localhost:9930?database=gorm"
func SQLServer(dsn string) gorm.Dialector {
	return sqlserver.Open(dsn)
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
func Postgres(dsn string) gorm.Dialector {
	return postgres.Open(dsn)
}

// dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
func ClickHouse(dsn string) gorm.Dialector {
	return clickhouse.Open(dsn)
}

var GormDriver map[string]func(string) gorm.Dialector = map[string]func(string) gorm.Dialector{
	"SQLITE":     Sqlite,
	"MYSQL":      MySQL,
	"SQLSERVER":  SQLServer,
	"POSTGRES":   Postgres,
	"CLICKHOUSE": ClickHouse,
}
