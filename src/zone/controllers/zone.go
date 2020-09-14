package controllers

import (
	"Gozone/library/enum"
	"Gozone/src/zone/models"
	"time"
)

type ZoneController struct {
	BaseHandler
}

func (this *ZoneController) Home() {

	sortType, _ := this.GetInt64("sortType", 0)
	//获取首页文章
	Articles, count, err := models.ArticleInstance.PageList(this.Pager.Offset, this.Pager.Limit, sortType)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	for k, v := range Articles {
		article, _ := models.ArticleClassInstance.FindArticleName(v.ArticleClass)
		Articles[k].ArticleClassName = article.ClassName
		Articles[k].CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
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
	this.Data["HomeContent"] = homeContent
	this.TplName = "base.html"
}

func (this *ZoneController) ArticleClass() {

	sortType, _ := this.GetInt64("sortType", 1)
	classId, _ := this.GetInt64("class_id", 0)
	//获取首页文章
	Articles, count, err := models.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, classId)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	for k, v := range Articles {
		article, _ := models.ArticleClassInstance.FindArticleName(v.ArticleClass)
		Articles[k].ArticleClassName = article.ClassName
		Articles[k].CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
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
	this.Data["HomeContent"] = homeContent
	this.TplName = "base.html"
}

func (this *ZoneController) ArticleTag() {
	this.TplName = "base.html"
}

func (this *ZoneController) Login() {
	this.TplName = "login.html"
}

func (this *ZoneController) Register() {
	this.TplName = "register.html"
}
