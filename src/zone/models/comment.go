package models

import "Gozone/library/conn"

type Comment struct {
	ID            int64         `json:"id" gorm:"column:id"`
	ArticleID     int64         `json:"article_id" gorm:"column:article_id"`
	Content       string        `json:"content" gorm:"column:content"`
	Floor         int64         `json:"floor" gorm:"column:floor"`
	CreateTime    int64         `json:"create_time" gorm:"column:create_time"`
	CreateTimeStr string        `json:"create_time_str"`
	UserId        int64         `json:"user_id"`
	UserInfo      *User         `json:"user_info"`
	CommentReply  *CommentReply `json:"comment_reply"`
}

func (this *Comment) TableName() string {
	return "comment"
}

func (this *Comment) GetArticleComment(articleId int64) (data []*Comment, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("article_id=?", articleId).Order("floor desc").Find(&data).Error
	return
}
