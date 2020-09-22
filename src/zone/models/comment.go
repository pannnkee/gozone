package models

import "Gozone/library/conn"

type Comment struct {
	ID         int64  `json:"id" gorm:"column:id"`
	ArticleID  int64  `json:"article_id" gorm:"column:article_id"` // 评论文章Id
	Content    string `json:"content" gorm:"column:content"`       // 评论
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	ParentID   int64  `json:"parent_id" gorm:"column:parent_id"` // 父级评论ID
	FromUid    int64  `json:"from_uid" gorm:"column:from_uid"`   // A-->B
	ToUid      int64  `json:"to_uid" gorm:"column:to_uid"`       // 如果没有to_id 说明是一级评论

	Floor    int64 `json:"floor"` //楼层
	FromUser *User	`json:"from_user"`
	ToUer    *User	`json:"to_user"`
}

func (this *Comment) TableName() string {
	return "comment"
}

func (this *Comment) GetFirstComment(articleId int64) (data []*Comment, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("create_time desc").Where("article_id=?", articleId).Where("to_uid=0").Find(&data).Error
	return
}
