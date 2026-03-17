package models

type ArticleCate struct {
	Id      int
	Title   string
	State   int
	Article []Article `gorm:"foreignkey:ArticleCateId"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
