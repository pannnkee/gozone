package models

// 关于
type About struct {
	ID      int64  `json:"id" gorm:"column:id"`
	Content string `json:"content" gorm:"column:content"`
}

func (this *About) TableName() string {
	return "about"
}
