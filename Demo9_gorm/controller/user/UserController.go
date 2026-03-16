package user

import (
	"Golang-learn/Demo9_gorm/models"
	"net/http"
	"time"

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

// CreateUser 新增用户
func CreateUser(c *gin.Context) {
	// 模拟新增用户
	user := models.User{
		Username: "李四",
		Age:      20,
		AddTime:  int(time.Now().Unix()),
	}
	// 新增用户到数据库
	models.DB.Create(&user)
	c.JSON(200, gin.H{
		"msg": "create user success",
	})
}

// ShowUser 查询所有用户
func ShowUser(c *gin.Context) {
	// 查询数据库
	var userList []models.User

	models.DB.Find(&userList)

	// 筛选年龄大于20的
	//models.DB.Where("age > ?", 20).Find(&userList)

	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

// EditUser 更新用户
func EditUser(c *gin.Context) {
	// 获取用户
	user := models.User{Id: 2}
	models.DB.Find(&user)
	// 设置更新后的用户信息
	user.Username = "王五"
	models.DB.Save(&user)
	// 更新用户
	models.DB.Updates(&user).Where("id=?", 2)
	c.JSON(200, gin.H{
		"msg": "edit user success",
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	// 先获取用户id
	user := models.User{Id: 3}
	// 只需要id即可实现删除
	models.DB.Delete(&user)

	c.JSON(200, gin.H{
		"msg": "delete user success",
	})
}
