package models

import "Gozone/library/conn"

type Comment struct {
	ID                   int64  `json:"id" gorm:"column:id"`                                           // 评论id
	UserID               int64  `json:"user_id" gorm:"column:user_id"`                                 // 评论人userId
	UserName             string `json:"user_name" gorm:"column:user_name"`                             // 评论人名称
	UserAvatar           string `json:"user_avatar" gorm:"-"`                                          //评论头像
	ArticleID            int64  `json:"article_id" gorm:"column:article_id"`                           // 评论的文章id
	ArticleTitle         string `json:"article_title" gorm:"column:article_title"`                     // 评论的文章标题
	ParentCommentID      int64  `json:"parent_comment_id" gorm:"column:parent_comment_id"`             // 父评论id
	ParentCommentUserID  int64  `json:"parent_comment_user_id" gorm:"column:parent_comment_user_id"`   // 父评论的用户id
	ReplyCommentID       int64  `json:"reply_comment_id" gorm:"column:reply_comment_id"`               // 被回复的评论id
	ReplyCommentUserID   int64  `json:"reply_comment_user_id" gorm:"column:reply_comment_user_id"`     // 被回复的评论用户id
	ReplyCommentUserName string `json:"reply_comment_user_name" gorm:"column:reply_comment_user_name"` //被评论用户名称
	CommentLevel         int64  `json:"comment_level" gorm:"column:comment_level"`                     // 评论等级[ 1 一级评论 默认 ，2 二级评论]
	Content              string `json:"content" gorm:"column:content"`                                 // 评论的内容
	Status               int64  `json:"status" gorm:"column:status"`                                   // 状态 (1 有效，0 逻辑删除)
	PraiseNum            int64  `json:"praise_num" gorm:"column:praise_num"`                           // 点赞数
	TopStatus            int64  `json:"top_status" gorm:"column:top_status"`                           // 置顶状态[ 1 置顶，0 不置顶 默认 ]
	CreateTime           int64  `json:"create_time" gorm:"column:create_time"`                         // 创建时间
	CreateTimeStr        string `json:"create_time_str" gorm:"column:create_time_str"`                 // 创建时间
	Floor                int64  `json:"floor" gorm:"-"`                                                //楼层

	SecondComment []*Comment `json:"second_comment"` //二级评论
}

func (this *Comment) TableName() string {
	return "comment"
}

func (this *Comment) AddComment() (err error) {
	db := conn.GetORMByName("zone")
	return db.Model(this).Create(&this).Error
}

func (this *Comment) GetFirstComment(articleId int64) (data []*Comment, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("create_time desc").Where("article_id=?", articleId).Where("comment_level=1").Find(&data).Error
	return
}

func (this *Comment) GetSecondComment(articleId, parentCommentId int64) (data []*Comment, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("create_time asc").Where("article_id=?", articleId).Where("comment_level=2").
		Where("parent_comment_id=? or reply_comment_id=?", parentCommentId, parentCommentId).Find(&data).Error
	return
}

func (this *Comment) GetCommentNumsAndHuman(articleId int64) (commentNums, Humans int64) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	db.Where("article_id=?", articleId).Count(&commentNums)
	db.Where("article_id=?", articleId).Group("user_id").Count(&Humans)
	return
}
