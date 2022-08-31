package database_test

import (
	"testing"

	"github.com/lovelacelee/clsgo/pkg/database"
	"github.com/lovelacelee/clsgo/pkg/log"
	// "github.com/lovelacelee/clsgo/pkg/utils"
)

func Test(t *testing.T) {
	Example()
}

func Example() {
	ExampleNew()
}

type User struct {
	database.Model
	Name string `gorm:"name"`
}

func ExampleNew() {
	db := database.New()
	if db != nil {
		defer db.Close()
		if db.Valid() {
			// db.Orm.Use(database.TracePlugin)
			db.Orm.AutoMigrate(&User{})
			user := User{
				Name: "Lee",
			}
			db.Orm.Create(&user)
			quser := User{}
			var count int64 = 0
			// db.Orm.Find(&quser, "id = ?", 2)
			db.CacheFind("", &quser, "id = ?", 2)
			db.Orm.Model(&User{}).Count(&count)
			log.Info(quser)
			log.Info(count)
		}
	}
}
