package models

// 文章
type Article struct {
	Id               int64  `gorm:"column:id" json:"id"`
	ArticleTitle     string `gorm:"column:article_title" json:"article_title"`
	ArticleClass     int64  `gorm:"column:article_class" json:"article_class"`
	ArticleClassName string `gorm:"-" json:"article_class_name"`
	SimpleContent    string `gorm:"column:simple_content" json:"simple_content"`
	Views            int64  `gorm:"column:views" json:"views"`
	CommentNumber    int64  `gorm:"column:comment_number" json:"comment_number"`
	Author           string `gorm:"column:author" json:"author"`
	Pic              string `gorm:"column:pic" json:"pic"`
	Carousel         int64  `gorm:"column:carousel" json:"carousel"`
	CreateTime       int64  `gorm:"column:create_time" json:"create_time"`
	CreatedTimeStr   string `gorm:"column:created_time_str" json:"create_time_str"`
	UpdateTime       int64  `gorm:"column:update_time" json:"update_time"`
	UpdateTimeStr    string `gorm:"column:update_time_str" json:"update_time_str"`
}

// 文章列表详情
type ArticleListResp struct {
	Article
	ArticleContent     string     `json:"article_content"`
	ArticleClassName   string     `json:"article_class_name"`
	ArticleContentNums int64      `json:"article_content_nums"`
	ArticleHumans      int64      `json:"article_humans"`
	ArticleTags        []*Tag     `json:"article_tags"`
	Emoji              [][]*Emoji `json:"emoji"`
	Comment            []*Comment `json:"comment"`
}

func (this *Article) TableName() string {
	return "article"
}
