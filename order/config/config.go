package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("order")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件出错，配置文件名：order,错误信息：%s", err)
	}

}
