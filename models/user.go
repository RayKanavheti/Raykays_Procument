package models

import (
	"github.com/jinzhu/gorm"
)

//User that falls part into an user
type User struct {
	gorm.Model
	UUID        string    `json:"-"`
	Username    string    `sql:"not null;unique" valid:"Required"`
	Email       string    `sql:"type:VARCHAR(50);not null;unique"`
	Password    string    `json:"-"`
	UserGroupID uint      `valid:"Required"`
	UserGroup   UserGroup `gorm:"save_associations:false"`
	Active      bool      `sql:"DEFAULT:1"`
	CellPhone   string    `sql:"type:VARCHAR(20)"`
	ResetKey    string    `json:"-"`
	Verified    bool
	Supplier    Supplier `gorm:"foreignkey:UserID"`
}
