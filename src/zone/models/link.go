package models

import "Gozone/library/conn"

type Link struct {
	Id      int64  `gorm:"column:id" json:"id"`
	Domain  string `gorm:"column:domain" json:"domain"`
	URL     string `gorm:"column:url" json:"url"`
	Content string `gorm:"column:content" json:"content"`
}

func (this *Link) TableName() string {
	return "link"
}

func (this *Link) FindLinks() (data []Link, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Find(&data).Error
	return
}

func (this *Link) GetAllData() (data []*Link, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("id asc").Find(&data).Error
	return
}
