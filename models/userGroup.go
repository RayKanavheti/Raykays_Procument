package models

import (
	"github.com/jinzhu/gorm"
)

//UserGroup for users
type UserGroup struct {
	gorm.Model
	Type   string
	Active bool `sql:"DEFAULT:1"`
}
