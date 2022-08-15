package db_test

import (
	"testing"

	"github.com/lovelacelee/clsgo/pkg/db"
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
	db.Model
	Name string `orm:"name"`
}

func ExampleNew() {
	db := db.New()
	if db != nil {
		defer db.Close()
		db.Orm.AutoMigrate(&User{})
		user := User{
			Name: "Lee",
		}
		db.Orm.Create(&user)
		db.Orm.Find(&user, "id = ?", 10)
		log.Info(user)
	}
}
