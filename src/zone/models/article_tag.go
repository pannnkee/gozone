package models

import "Gozone/library/conn"

type ArticleTag struct {
	Id        int64 `gorm:"column:id" json:"id"`
	ArticleId int64 `gorm:"column:article_id" json:"article_id"`
	TagId     int64 `gorm:"column:tag_id" json:"tag_id"`
}

func (this *ArticleTag) TableName() string {
	return "article_tag"
}

func (this *ArticleTag) FindTags(articleId int64) (data []*ArticleTag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("article_id=?", articleId).Find(&data).Error
	return
}

func (this *ArticleTag) FindArticles(tagId int64) (data []*ArticleTag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("tag_id=?", tagId).Find(&data).Error
	return
}
