package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

// User is just a sample type
type User struct {
	ID   string `json:"id" description:"identifier of the user" gorm:"primary_key"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
	Dogs *[]Dog
}

type Dog struct {
	gorm.Model
	Name     *string
	Color    []sql.NullString
	BirthDay *time.Time
	NoJSON   string `json:"-"`
}
