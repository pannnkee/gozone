package models

import "Gozone/library/conn"

type Tag struct {
	Id      int    `gorm:"column:id" json:"id"`
	TagName string `gorm:"column:tag_name" json:"tag_name"`
}

func (this *Tag) TableName() string {
	return "tag"
}

func (this *Tag) GetTagName(id int64) (name string, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("id=?", id).Find(&this).Error
	name = this.TagName
	return
}
