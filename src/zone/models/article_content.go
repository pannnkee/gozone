package models

import "Gozone/library/conn"

type ArticleContent struct {
	Id      string `gorm:"column:id" json:"id"`
	Content string `gorm:"column:content" json:"content"`
}

func (this *ArticleContent) TableName() string {
	return "article_content"
}

func (this *ArticleContent) Get(id int64) (err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("id=?", id).Take(&this).Error
	return
}
