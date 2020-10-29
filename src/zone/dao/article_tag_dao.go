package dao

import (
	"Gozone/library/conn"
	"Gozone/src/zone/models"
)

type ArticleTagDao struct {}

// 根据文章ID获取所有tag
// @parma articleId 文章Id
// @return data 文章所有tag
// @return err 错误信息
func (this *ArticleTagDao) FindTags(articleId int64) (data []*models.ArticleTag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(ArticleTagInstance)
	err = db.Where("article_id=?", articleId).Find(&data).Error
	return
}

// 根据tagId获取所有包含此tag的数据
// @param tagId 文章tagId
// @return data 包含此tag的所有数据
// @return err 错误信息
func (this *ArticleTagDao) FindArticles(tagId int64) (data []*models.ArticleTag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(ArticleTagInstance)
	err = db.Where("tag_id=?", tagId).Find(&data).Error
	return
}
