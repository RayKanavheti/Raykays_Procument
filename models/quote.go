package models

import (
	"github.com/jinzhu/gorm"
)

//Quote for Orders
type Quote struct {
	gorm.Model
	Active      bool `sql:"DEFAULT:1"`
  QouteProducts []*QuoteProduct `gorm:"foreignkey:QouteID"`
	TotalAmount float32
	IsAccepted  bool
	SupplierID  uint
	OrderID     uint
}
