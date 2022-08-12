package db_test

import (
	"testing"

	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/db"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func Test(t *testing.T) {
	Example()
}

func Example() {

	ExampleNew()
	ExampleNewSqlite()
}

type User struct {
	db.Model
	Name string `orm:"name"`
}

func ExampleNew() {
	db := db.New()
	m := db.Model(User{})
	log.Infoi(m.Data(clsgo.Map{"name": "john"}).Insert())
}

func ExampleNewSqlite() {
	db := db.NewSqlite("file::memory:?cache=shared")
	db.Instance.AutoMigrate(&User{})
	user := User{
		Name: "Lee",
	}
	db.Instance.Create(&user)
	db.Instance.Find(&user, "id = ?", 10)
	log.Info(user)
}
