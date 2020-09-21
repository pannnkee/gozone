package models

import "Gozone/library/conn"

type CommentReply struct {
	Id        int    `json:"id" gorm:"column:id"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id"` // 表示该回复挂在的根评论 id
	ReplyId   int    `json:"reply_id" gorm:"column:reply_id"`
	Content   string `json:"content" gorm:"column:content"`   // 回复内容
	FromUid   int    `json:"from_uid" gorm:"column:from_uid"` // 回复用户id
	ToUid     int    `json:"to_uid" gorm:"column:to_uid"`     // 目标用户 id
	UserInfo  *User  `json:"user_info"`
}

func (this *CommentReply) TableName() string {
	return "comment_reply"
}

func (this *CommentReply) GetReplyByCommentId(commentId int64) (data []*CommentReply, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("reply_id desc").Where("comment_id=?", commentId).Find(&data).Error
	return
}
