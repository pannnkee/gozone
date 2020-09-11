package models

import "Gozone/library/conn"

type ArticleClass struct {
	Id        int64  `gorm:"column:id" json:"id"`
	ClassName string `gorm:"column:class_name" json:"class_name"`
	Nums      int64  `gorm:"-" json:"nums"`
}

func (this *ArticleClass) TableName() string {
	return "article_class"
}

func (this *ArticleClass) FindArticleName(id int64) (data *ArticleClass, err error) {
	db := conn.GetORMByName("zone")
	data = new(ArticleClass)
	err = db.Where("id=?", id).First(&data).Error
	return
}

func (this *ArticleClass) FindAllArticleClass() (data []*ArticleClass, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Find(&data).Error
	return
}
