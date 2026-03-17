package controller

import (
	"Golang-learn/Demo11_transaction/models"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (UserController) UserInfo(c *gin.Context) {

	// 开启事务
	tx := models.DB.Begin()

	// id为1的人age涨1岁
	user := models.User{Id: 1}
	tx.Find(&user)
	user.Age += 1
	tx.Save(&user)

	// 触发异常，此时程序会继续执行，如果不显式调用Rollback()，则会运行后面的Commit，导致错误的数据被提交
	if true {
		tx.Rollback()
		logger.Error("error")
	}

	// 触发错误
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if true {
		panic("panic")
	}
	// id为2的人age涨1岁
	user = models.User{Id: 2}
	tx.Find(&user)
	user.Age += 1
	tx.Save(&user)

	// 事务提交
	tx.Commit()

	c.JSON(200, gin.H{
		"message": "user info",
	})
}
