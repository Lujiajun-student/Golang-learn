package models

import (
	"github.com/bytedance/gopkg/util/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitMySQL() {
	dsn := "root:jia.1113.me.@tcp(127.0.0.1:3307)/gin?charset-utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}
}
