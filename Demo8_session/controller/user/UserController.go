package user

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	// 取出session对象
	session := sessions.Default(c)
	// 使用session对象来设置数据
	session.Set("username", "张三")
	// 设置完数据后，需要调用session.Save()方法来保存数据
	err := session.Save()
	if err != nil {
		logger.Warn("session save failed, err: %v", err)
	}
	c.JSON(200, gin.H{
		"msg": "user info",
	})
}

func UserNews(c *gin.Context) {
	// 取出session对象
	session := sessions.Default(c)
	// 通过session对象来获取数据
	username := session.Get("username")
	c.JSON(200, gin.H{
		"msg":      "user news",
		"username": username,
	})
}
