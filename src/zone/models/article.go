package models

import "Gozone/library/conn"

type Article struct {
	Id            int64  `gorm:"column:id" json:"id"`
	ArticleTitle  string `gorm:"column:article_title" json:"article_title"`
	ArticleClass  int64  `grom:"column:article_class" json:"article_class"`
	Views         int64  `gorm:"column:views" json:"views"`
	CommentNumber int64  `gorm:"column:comment_number" json:"comment_number"`
	CreateTime    int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime    int64  `gorm:"column:update_time" json:"update_time"`
}

func (this *Article) TableName() string {
	return "article"
}

func (this *Article) PageList(offset, limit int64) (datas []*Article, count int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Offset(offset).Limit(limit).Order("id desc").Find(&datas).Error
	err = db.Count(&count).Error
	return
}
