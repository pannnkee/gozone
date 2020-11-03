package dao

import (
	"gozone/library/conn"
	"gozone/src/zone/models"
)

type EmojiDao struct {}

// 获取所有Emoji
// @return data 所有Emoji
func (this *EmojiDao) GetAll() (data []*models.Emoji, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.EmojiInstance)
	err = db.Order("id asc").Find(&data).Error
	return
}
