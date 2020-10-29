package models

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

