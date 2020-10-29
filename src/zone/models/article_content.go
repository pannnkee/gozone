package models

type ArticleContent struct {
	Id      string `gorm:"column:id" json:"id"`
	Content string `gorm:"column:content" json:"content"`
}

func (this *ArticleContent) TableName() string {
	return "article_content"
}
