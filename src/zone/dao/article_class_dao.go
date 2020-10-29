package dao

import (
	"Gozone/library/conn"
	"Gozone/src/zone/models"
)

type ArticleClassDao struct {}

// 根据文章分类ID获取分类消息
// @param id 文章分类ID
// @return data 文章分类信息
// @return err 错误信息
func (this *ArticleClassDao) Get(id int64) (data models.ArticleClass, err error) {
	db := conn.GetORMByName("zone")
	err = db.Where("id=?", id).Take(&data).Error
	return
}

// 获取所有文章分类
// @return data 所有文章分类信息
// @return err 错误信息
func (this *ArticleClassDao) GetAll() (data []*models.ArticleClass, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleClassInstance)
	err = db.Order("id asc").Find(&data).Error
	return
}
