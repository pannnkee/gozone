package dao

import (
	"Gozone/library/conn"
	"Gozone/src/zone/models"
)

type AboutDao struct {}

// 获取关于所有信息
// @return data 关于所有信息
// @return err 错误信息
func (this *AboutDao) GetAllData() (data []*models.About, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.AboutInstance)
	err = db.Find(&data).Error
	return
}
