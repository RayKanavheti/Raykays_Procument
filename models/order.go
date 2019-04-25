package models

import (
	"github.com/jinzhu/gorm"
)

//Order for users
type Order struct {
	gorm.Model
	Active     bool `sql:"DEFAULT:1"`
  OrderProducts []*OrderProduct `gorm:"foreignkey:OrderID"`
  Quotes []*Quote `gorm:"foreignkey:OrderID"`
	SupplierID uint
}
