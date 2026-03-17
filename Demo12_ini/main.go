package main

import (
	"Golang-learn/Demo12_ini/models"
	"Golang-learn/Demo12_ini/router"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {
	r := gin.Default()
	models.InitMySQL()
	router.InitUserRouter(r)
	config, err := ini.Load("./Demo12_ini/config/app.ini")
	if err != nil {
		logger.Error(err.Error())
	}
	config.Section("").Key("set_new_key").SetValue("new_value")
	// 写完后需要保存
	err = config.SaveTo("./Demo12_ini/config/app.ini")
	if err != nil {
		logger.Error(err.Error())
	}
	r.Run(":8080")
}
