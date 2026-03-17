package models

import (
	"fmt"

	"github.com/bytedance/gopkg/util/logger"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitMySQL() {
	config, err := ini.Load("./Demo12_ini/config/app.ini")
	if err != nil {
		logger.Error(err.Error())
	}
	username := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	database := config.Section("mysql").Key("database").String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, ip, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}
}
