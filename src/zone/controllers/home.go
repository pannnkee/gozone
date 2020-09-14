package controllers

import (
	"Gozone/library/enum"
	"Gozone/src/zone/models"
	"time"
)

type HomeController struct {
	BaseHandler
}

func (this *HomeController) Content() {

	sortType, _ := this.GetInt64("sortType", 1)

	//获取首页文章
	Articles, count, err := models.ArticleInstance.PageList(this.Pager.Offset, this.Pager.Limit, sortType)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	for _, v := range Articles {
		article, _ := models.ArticleClassInstance.FindArticleName(v.ArticleClass)
		v.ArticleClassName = article.ClassName
		v.CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
	}

	this.Pager.Count = count

	//获取首页标签
	tag, err := models.TagInstance.GetAllTag()
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	//获取首页文章分类
	class, err := models.ArticleClassInstance.FindAllArticleClass()
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	for _, v := range class {
		nums, _ := models.ArticleInstance.FindClassNums(v.Id)
		v.Nums = nums
	}


	//获取友情链接
	links, err := models.LinkInstance.FindLinks()
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	homeContent := new(models.HomeContent)
	homeContent.Articles = Articles
	homeContent.Tags = tag
	homeContent.ArticleClass = class
	homeContent.Links = links
	homeContent.PannnkeeZone = "Pannnkee's Zone"
	homeContent.SortType = sortType
	homeContent.ContentType = enum.DefaultType

	this.Response(enum.DefaultSuccess, "", homeContent)
}

func (this *HomeController) ContentClass() {

}

func (this *HomeController) ContentTag() {

}


