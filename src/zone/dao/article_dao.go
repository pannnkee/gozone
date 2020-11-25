package dao

import (
	"gozone/library/conn"
	"gozone/library/enum"
	"gozone/src/zone/models"
)

type ArticleDao struct {}

// 文章分页列表
// @param offset 偏移量
// @param limit 限制
// @param sortType 排序方式
// @return data 文章列表数据
// @return count 文章数量
// @return err 错误信息
func (this *ArticleDao) PageList(offset, limit, sortType int64) (data []models.Article, count int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)

	if sortType == int64(enum.HotSort) {
		err = db.Offset(offset).Limit(limit).Order("views desc").Find(&data).Error
	} else {
		err = db.Offset(offset).Limit(limit).Order("create_time asc").Find(&data).Error
	}
	err = db.Count(&count).Error
	return
}

// 获取分类下文章列表
// @param offset 偏移量
// @param limit 限制
// @param sortType 排序方式
// @param contentType 文章分类
// @return data 文章列表
// @return count 数量
// @return err 错误信息
func (this *ArticleDao) PageListClass(offset, limit, sortType, contentType int64) (data []*models.Article, count int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)

	if contentType > 0 {
		db = db.Where("article_class=?", contentType)
	}
	if sortType == int64(enum.HotSort) {
		err = db.Offset(offset).Limit(limit).Order("views desc").Find(&data).Error
	} else {
		err = db.Offset(offset).Limit(limit).Order("create_time desc").Find(&data).Error
	}
	err = db.Count(&count).Error
	return
}

// 获取轮播图文章
func (this *ArticleDao) GetCarouselArticle() (data []*models.Article, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)
	err = db.Where("carousel=1").Find(&data).Error
	return
}

// 根据文章ID获取文章详细信息
// @param id 文章id
// @return err 错误信息
func (this *ArticleDao) Get(id int64) (article models.Article, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(&models.ArticleInstance)
	err = db.Where("id=?", id).Find(&article).Error
	return
}

// 更新文章观看次数
// @param id 文章id
// @return err 错误信息
func (this *ArticleDao) UpdateViews(id int64) (err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)

	var article models.Article
	err = db.First(&article, id).Error
	if err != nil {
		return err
	}
	return db.Model(models.ArticleInstance).Where("id=?", id).Update("views", article.Views+1).Error
}

// 查找一个分类下文章数量
// @param classId 分类Id
// @return nums 数量
// @return err 错误信息
func (this *ArticleDao) FindClassNums(classId int64) (nums int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)
	err = db.Where("article_class=?", classId).Count(&nums).Error
	return
}

// 根据文章id 查找所有文章
// @param id 文章id合集
// @return data 文章合集
// @return err 错误信息
func (this *ArticleDao) FindArticles(id []int64) (data []*models.Article, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)
	db = db.Where("id in (?)", id)
	err = db.Find(&data).Error
	return
}

// 获取所有文章信息
// @return data 所有文章数据
// @return err  错误信息
func (this *ArticleDao) GetAllData() (data []*models.Article, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)
	err = db.Order("id asc").Find(&data).Error
	return
}

// 根据分类ID获取该分类下所有文章
// @param classID 分类ID
// @return count 数量
// @return err 错误信息
func (this *ArticleDao) GetArticleClassNums(classID int64) (count int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)
	err = db.Where("article_class=?", classID).Count(&count).Error
	return
}

// 获取所有文章数量
func (this *ArticleDao) GetArticleNums() (count int64, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.ArticleInstance)
	err = db.Count(&count).Error
	return
}
