package dao

import (
	"Gozone/library/conn"
	"Gozone/src/zone/models"
)

type TimelineDao struct {}

// 获取TimeLine所有信息
// @return data 所有数据
// @return err 错误信息
func (this *TimelineDao) GetAll() (data []*models.Timeline, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.TimeLineInstance)
	err = db.Find(&data).Error
	return
}
