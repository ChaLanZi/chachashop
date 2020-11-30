package models

import "github.com/jinzhu/gorm"

// Discount 订单折扣表
type Discount struct {
	gorm.Model
	Price float64
}
