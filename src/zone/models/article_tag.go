package models

// 文章tag
type ArticleTag struct {
	Id        int64 `gorm:"column:id" json:"id"`
	ArticleId int64 `gorm:"column:article_id" json:"article_id"`
	TagId     int64 `gorm:"column:tag_id" json:"tag_id"`
}

func (this *ArticleTag) TableName() string {
	return "article_tag"
}
