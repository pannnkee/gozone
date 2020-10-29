package dao

import (
	"Gozone/library/conn"
	"Gozone/src/zone/models"
)

type LinkDao struct {}

// 获取所有链接
// @return data 所有链接
// @return err 错误信息
func (this *LinkDao) GetAll() (data []*models.Link, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("id asc").Find(&data).Error
	return
}
