// Table use gorm's migrate speciality

package db

import (
	"gorm.io/gorm"
)

type Table struct {
	Db gorm.DB
}
