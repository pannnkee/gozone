package model_view

type CommentWeb struct {
	RepID         int64  `json:"rep_id"`
	Content       string `json:"content"`
	ArticleId     int64  `json:"article_id"`
	ReplyFatherID int64  `json:"reply_father_id"` //回复二级评论 需要定位是在哪一个一级评论
	RepUserID     int64  `json:"rep_user_id"` //我回复 楼层用户的ID
}
