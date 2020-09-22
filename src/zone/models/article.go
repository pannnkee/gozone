package models

import (
	"Gozone/library/conn"
	"Gozone/library/enum"
)

type Article struct {
	Id               int64  `gorm:"column:id" json:"id"`
	ArticleTitle     string `gorm:"column:article_title" json:"article_title"`
	ArticleClass     int64  `gorm:"column:article_class" json:"article_class"`
	ArticleClassName string `gorm:"column:article_class_name" json:"article_class_name"`
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

// 文章详情 Text
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

func (this *Article) PageList(offset, limit, sortType int64) (data []Article, count int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)

	if sortType == int64(enum.HotSort) {
		err = db.Offset(offset).Limit(limit).Order("views desc").Find(&data).Error
	} else {
		err = db.Offset(offset).Limit(limit).Order("create_time asc").Find(&data).Error
	}
	err = db.Count(&count).Error
	return
}

// 获取分类下文章列表
func (this *Article) PageListClass(offset, limit, sortType, contentType int64) (data []Article, count int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)

	if contentType > 0 {
		db = db.Where("article_class=?", contentType)
	}
	if sortType == int64(enum.HotSort) {
		err = db.Offset(offset).Limit(limit).Order("views desc").Find(&data).Error
	} else {
		err = db.Offset(offset).Limit(limit).Order("create_time asc").Find(&data).Error
	}
	err = db.Count(&count).Error
	return
}

func (this *Article) Get(id int64) (err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("id=?", id).Find(&this).Error
	return
}

func (this *Article) FindClassNums(classId int64) (nums int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("article_class=?", classId).Count(&nums).Error
	return
}

func (this *Article) FindArticles(id []int64) (data []Article, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	db = db.Where("id in (?)", id)
	err = db.Find(&data).Error
	return
}

func (this *Article) GetAllData() (data []*Article, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Order("id asc").Find(&data).Error
	return
}
