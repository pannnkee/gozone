package models

import "Gozone/library/conn"

type About struct {
	ID      int64    `json:"id" gorm:"column:id"`
	Content string `json:"content" gorm:"column:content"`
}

func (this *About) TableName() string {
	return "about"
}

func (this *About) GetAllData() (data []*About, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Find(&data).Error
	return
}