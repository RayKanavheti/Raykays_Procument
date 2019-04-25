package models

import (
	"github.com/jinzhu/gorm"
)

//QouteProduct for a quote
type QuoteProduct struct {
	gorm.Model
	Active      bool   `sql:"DEFAULT:1"`
	Title       string `sql:"type:VARCHAR(100)"`
	Description string `sql:"type:VARCHAR(200)"`
	CostEach    float32
	Quantity    float32
	TotalCost   float32
	QouteID     uint
}
