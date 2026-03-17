package models

type Article struct {
	Id            int
	Title         string
	ArticleCateId int
	State         int
	ArticleCate   ArticleCate
}

func (Article) TableName() string {
	return "article"
}
