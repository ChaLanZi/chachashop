package models

import "time"

// Order 订单主表
type Order struct {
	OrderId    string `json:"order_id"`
	Goods      []Goods
	Payments   []Payment
	Discounts  []Discount
	Version    uint32
	CreateTime time.Time
	Note       string
}
