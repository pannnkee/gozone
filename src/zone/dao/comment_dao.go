package dao

import (
	"Gozone/library/conn"
	"Gozone/src/zone/models"
)

type CommentDao struct {}

// 添加一条评论
// @param comment 评论内容
// @return err 错误信息
func (this *CommentDao) AddComment(comment *models.Comment) (err error) {
	db := conn.GetORMByName("zone")
	return db.Model(comment).Create(&comment).Error
}

// 获取一级评论
// @param articleId 文章Id
// @return data 所有一级评论
// @return err 错误信息
func (this *CommentDao) GetFirstComment(articleId int64) (data []*models.Comment, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(&models.CommentInstance)
	err = db.Order("create_time desc").Where("article_id=?", articleId).Where("comment_level=1").Find(&data).Error
	return
}

// 获取二级评论
// @param articleId 文章Id
// @param parentCommentId 父评论Id
// @return data 所有二级评论
// @return err 错误信息
func (this *CommentDao) GetSecondComment(articleId, parentCommentId int64) (data []*models.Comment, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("create_time asc").Where("article_id=?", articleId).Where("comment_level=2").
		Where("parent_comment_id=? or reply_comment_id=?", parentCommentId, parentCommentId).Find(&data).Error
	return
}

// 获取评论数、参与人数
// @param articleId 文章Id
// @return commentNums 文章评论数
// @return Humans 参与人数
func (this *CommentDao) GetCommentNumsAndHuman(articleId int64) (commentNums, Humans int64) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	db.Where("article_id=?", articleId).Count(&commentNums)
	db.Where("article_id=?", articleId).Group("user_id").Count(&Humans)
	return
}
