package models

import (
	"Gozone/library/cache"
	"Gozone/library/conn"
	cache2 "Gozone/src/zone/cache"
	"Gozone/src/zone/models"
)

type Comment struct {
	Id         int64        `json:"id" gorm:"column:id"`
	ArticleID  int64        `json:"article_id" gorm:"column:article_id"` // 评论文章Id
	Content    string       `json:"content" gorm:"column:content"`       // 评论
	CreateTime int64        `json:"create_time" gorm:"column:create_time"`
	ParentID   int64        `json:"parent_id" gorm:"column:parent_id"` // 父级评论ID
	FromUid    int64        `json:"from_uid" gorm:"column:from_uid"`   // A-->B
	ToUid      int64        `json:"to_uid" gorm:"column:to_uid"`       // 如果没有to_id 说明是一级评论
	FromUser   *models.User `json:"from_user"`
	ToUser     *models.User `json:"to_user"`
}

func (this *Comment) TableName() string {
	return "comment"
}

func (this *Comment) GetFirstComment(articleId int64) (data []*Comment, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("create_time desc").Where("article_id=?", articleId).Where("to_uid=null").Find(&data).Error

	for _, v := range data {
		if v.FromUid > 0 {
			userInterface, _ := new(cache.Helper).GetByItemKey(new(cache2.UserCache), v.FromUid)
			v.FromUser = userInterface.(*models.User)
		}
	}

	return
}
