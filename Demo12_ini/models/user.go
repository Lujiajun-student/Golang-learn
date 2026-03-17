package models

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

func (User) TableName() string {
	return "user"
}
