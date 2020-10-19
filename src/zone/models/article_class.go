package models

import "Gozone/library/conn"

type ArticleClass struct {
	Id                int64  `gorm:"column:id" json:"id"`
	Url               string `gorm:"column:url" json:"url"`
	Nums              int64  `gorm:"-" json:"nums"`
	ClassName         string `gorm:"column:class_name" json:"class_name"`
	ClassIntroduction string `gorm:"column:class_introduction" json:"class_introduction"`
}

func (this *ArticleClass) TableName() string {
	return "article_class"
}

func (this *ArticleClass) Get(id int64) (data *ArticleClass, err error) {
	db := conn.GetORMByName("zone")
	data = new(ArticleClass)
	err = db.Where("id=?", id).First(&data).Error
	return
}

func (this *ArticleClass) FindAllArticleClass() (data []ArticleClass, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Find(&data).Error
	return
}

func (this *ArticleClass) GetAllData() (data []*ArticleClass, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("id asc").Find(&data).Error
	return
}

