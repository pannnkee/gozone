package models

import "Gozone/library/conn"

type Tag struct {
	Id         int64  `gorm:"column:id" json:"id"`
	TagName    string `gorm:"column:tag_name" json:"tag_name"`
	TagNum     int64  `gorm:"column:tag_num" json:"tag_num"`
	URL        string `gorm:"column:url" json:"url"`
	TagContent string `gorm:"column:tag_content" json:"tag_content"`
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

func (this *Tag) GetTag(id int64) (data Tag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("id=?", id).Take(&data).Error
	return
}

func (this *Tag) GetTags(id []int64) (data []*Tag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	db = db.Where("id in (?)", id)
	err = db.Find(&data).Error
	return
}

func (this *Tag) GetAllData() (data []*Tag, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("id asc").Find(&data).Error
	return
}
