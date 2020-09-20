package models

type Comment struct {
	ID        int    `json:"id" gorm:"column:id"`
	ArticleID int    `json:"article_id" gorm:"column:article_id"`
	UserID    int    `json:"user_id" gorm:"column:user_id"`
	Content   string `json:"content" gorm:"column:content"`
}

func (this *Comment) TableName() string {
	return "comment"
}
