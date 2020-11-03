package dao

import (
	"gozone/library/conn"
	"gozone/src/zone/models"
)

type ArticleContentDao struct {}

// 根据文章ID获取文章内容
// @param id 文章id
// @return articleContent 文章内容
// @return err 错误信息
func (this *ArticleContentDao) Get(id int64) (articleContent models.ArticleContent, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleContentInstance)
	err = db.Where("id=?", id).Take(&articleContent).Error
	return
}

// 获取所有文章内容
// @return data 所有文章内容
// @return err 错误信息
func (this *ArticleContentDao) GetAll() (data []*models.ArticleContent, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleContentInstance)
	err = db.Order("id asc").Find(&data).Error
	return
}

