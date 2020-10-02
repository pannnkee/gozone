package models

import "Gozone/library/conn"

type Timeline struct {
	ID            int64  `json:"id" gorm:"column:id"`
	Title         string `json:"title" gorm:"column:title"`
	Side          string `json:"side" gorm:"column:side"`
	Icon          string `json:"icon" gorm:"column:icon"`
	IconColor     string `json:"icon_color" gorm:"column:icon_color"`
	StatNum       int    `json:"stat_num" gorm:"column:stat_num"`
	Markdown      string `json:"markdown" gorm:"column:markdown"`
	UpdateTime    int64  `json:"update_time" gorm:"column:update_time"`
	UpdateTimeStr string `json:"update_time_str"`
}

func (this *Timeline) TableName() string {
	return "timeline"
}

func (this *Timeline) GetAllData() (data []*Timeline, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Find(&data).Error
	return
}
