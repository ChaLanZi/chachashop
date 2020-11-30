package models

import "github.com/jinzhu/gorm"

//Payment 订单支付表
type Payment struct {
	gorm.Model
	Total  float64
	Status uint
	Source uint
}
