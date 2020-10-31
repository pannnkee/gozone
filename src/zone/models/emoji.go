package models

// Emoji 表情
type Emoji struct {
	Id        int32  `gorm:"column:id" json:"id"`
	Src       string `gorm:"column:src" json:"src"`
	Title     string `gorm:"column:title" json:"title"`
	Alt       string `gorm:"column:alt" json:"alt"`
	DataEmoji string `gorm:"column:data_emoji" json:"data_emoji"`
}

func (this *Emoji) TableName() string {
	return "emoji"
}
