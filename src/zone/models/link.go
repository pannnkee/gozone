package models

// 友情链接
type Link struct {
	Id      int64  `gorm:"column:id" json:"id"`
	Domain  string `gorm:"column:domain" json:"domain"`
	URL     string `gorm:"column:url" json:"url"`
	Content string `gorm:"column:content" json:"content"`
}

func (this *Link) TableName() string {
	return "link"
}
