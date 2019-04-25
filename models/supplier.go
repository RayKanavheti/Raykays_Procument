package models

import (
	"github.com/jinzhu/gorm"
)

//Supplier that falls part into an user
type Supplier struct {
	gorm.Model
	RegisteredName   string `sql:"type:VARCHAR(200)"`
	RegisteredNumber string `sql:"type:VARCHAR(50)"`
	Address          string `sql:"type:VARCHAR(100)"`
	City             string `sql:"type:VARCHAR(30)"`
	Country          string `sql:"type:VARCHAR(30)"`
	Phone            string `sql:"type:VARCHAR(30)"`
	Email            string `sql:"type:VARCHAR(50);not null;unique"`
	UserID           uint
	Qoute            []*Quote `gorm:"foreignkey:SupplierID"`
}
