package models

import "github.com/jinzhu/gorm"

// Goods 订单商品表
type Goods struct {
	gorm.Model
	OrderId string
	Name    string `gorm:"type:varchar(255);not null"`
	Picture string `gorm:"type:varchar(100)"`
	Price   float64
	Weight  uint32
	Number  uint32
}
