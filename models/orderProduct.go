package models
import (
	"github.com/jinzhu/gorm"
)

//OrderProduct for an order
type OrderProduct struct {
	gorm.Model
	Active      bool   `sql:"DEFAULT:1"`
	Title       string `sql:"type:VARCHAR(100)"`
	Description string `sql:"type:VARCHAR(200)"`
	Quantity    float32
	OrderID     uint
}
