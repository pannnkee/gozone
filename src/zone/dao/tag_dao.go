package dao

import (
	"Gozone/library/conn"
	"Gozone/src/zone/models"
)

type TagDao struct {}

// 根据TagId 获取对应名称
// @param tagId
// @return name tag名称
// @return err 错误信息
func (this *TagDao) GetTagName(tagId int64) (name string, err error) {
	db := conn.GetORMByName("zone")

	Tag := new(models.Tag)
	db = db.Model(Tag)
	err = db.Where("id=?", tagId).Find(&this).Error
	name = Tag.TagName
	return
}

// 根据tagId获取tag详细信息
// @param tagId
// @return data tag详细信息
// @return err 错误信息
func (this *TagDao) Get(tagId int64) (data models.Tag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.TagInstance)
	err = db.Where("id=?", tagId).Take(&data).Error
	return
}

// 根据tagIds获取tag详细信息
// @param tagIds
// @return data tag详细信息
// @return err 错误信息
func (this *TagDao) GetTags(tagIds []int64) (data []*models.Tag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.TagInstance)
	db = db.Where("id in (?)", tagIds)
	err = db.Find(&data).Error
	return
}

// 获取所有Tag数据
// @return data 所有数据
// @return err 错误信息
func (this *TagDao) GetAll() (data []*models.Tag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.TagInstance)
	err = db.Order("id asc").Find(&data).Error
	return
}
