package models

// User 用户模型
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// TableName 接口，指定表名
func (User) TableName() string {
	return "user"
}
