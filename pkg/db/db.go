// Package Db wraps goframe gdb for convenience
package db

import (
	// "context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
	// GDB drivers[Choose right driver with config]
	_ "github.com/gogf/gf/contrib/drivers/clickhouse/v2" //clickhouse
	_ "github.com/gogf/gf/contrib/drivers/mssql/v2"      //sqlserver
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"      //mysql
	_ "github.com/gogf/gf/contrib/drivers/oracle/v2"     //oracle
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"      //pgsql
)

type Db = gdb.DB

func New() Db {
	instance, err := gdb.Instance("test")
	utils.WarnIfError(err, log.Errorf)
	return instance
}
