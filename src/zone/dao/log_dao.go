package dao

import (
	"gozone/library/conn"
	"gozone/src/zone/models"
)

type LogDao struct {}

// 添加一条登录日志
// @return err 错误信息
func (this *LogDao) AddLoginLog(loginLog *models.Log) (err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.LogInstance)
	return db.Create(&loginLog).Error
}
