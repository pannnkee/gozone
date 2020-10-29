package models

type ArticleClass struct {
	Id                int64  `gorm:"column:id" json:"id"`
	Nums              int64  `gorm:"-" json:"nums"`
	ClassName         string `gorm:"column:class_name" json:"class_name"`
	ClassIntroduction string `gorm:"column:class_introduction" json:"class_introduction"`
}

func (this *ArticleClass) TableName() string {
	return "article_class"
}


