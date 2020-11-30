package repo

import (
	"chachashop/order/repo/models"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func Open() {
	db := openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)

	setUpDB(db)

	DB = db

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "chacha_" + defaultTableName
	}

	models.Migrate(db)
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)?charset=utf8mb4&parseTime=%t%local=%s",
		username, password, addr, name, true, "Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		// 连接第三方库直接出错，打log直接处理，不向上抛出
		logrus.Fatal("数据库连接失败，数据库名字：%s,错误信息：%s", name, err)
	}
	return db
}

func setUpDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("logLevel"))
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	db.SingularTable(true)
}
