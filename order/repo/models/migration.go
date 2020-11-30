package models

import "github.com/jinzhu/gorm"

// Migrate 数据库自动迁移
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Goods{})
	db.AutoMigrate(&Payment{})
	db.AutoMigrate(&Discount{})
}
