package controllers

import (
	"Gozone/library/cache"
	"Gozone/library/enum"
	"Gozone/library/logger"
	cache2 "Gozone/src/zone/cache"
	"Gozone/src/zone/models"
	"fmt"
	"time"
)

type ZoneController struct {
	BaseHandler
}

func (this *ZoneController) Home() {

	cookie := this.Ctx.GetCookie("toggleTheme")
	fmt.Println(cookie)

	sortType, _ := this.GetInt64("sortType", 1)
	contentType, _ := this.GetInt64("contentType", 0)

	var Articles []models.Article
	var TopContent models.TopContent
	var count int64
	var err error

	homeContent := new(models.HomeContent)

	if enum.ContentType(contentType) == enum.DefaultType {
		//获取首页文章
		Articles, count, err = models.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
		if err != nil {
			this.Response(enum.DefaultError, err.Error())
			return
		}

		for k, v := range Articles {

			articleClassInterface, _ := new(cache.Helper).GetByItemKey(new(cache2.ArticleClassCache), v.ArticleClass)
			article := articleClassInterface.(*models.ArticleClass)

			Articles[k].ArticleClassName = article.ClassName
			Articles[k].CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
		}
		this.Pager.Count = count

		//获取轮播图

	} else {
		//获取base_top内容

		// 文章分类
		if enum.ContentType(contentType) < enum.Mysql {
			class, err := models.ArticleClassInstance.FindArticleClassName(contentType)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}
			Articles, _, err = models.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}

			for k, v := range Articles {
				Articles[k].CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
			}

			TopContent.ContentNum = class.Nums
			TopContent.ContentText = class.ClassIntroduction
			TopContent.TopContentClass = "文章分类"
			TopContent.TopContentName = class.ClassName
			TopContent.TopArticle = Articles
		} else {
			// 标签分类
			tag, err := models.TagInstance.GetTag(contentType - 100)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}

			var Tag[]int64
			articles, _ := models.ArticleTagInstance.FindArticles(tag.Id)
			for _, v := range articles {
				Tag = append(Tag, v.TagId)
			}

			TagArticles, err := models.ArticleInstance.FindArticles(Tag)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}
			TopContent.ContentNum = tag.TagNum
			TopContent.ContentText = tag.TagContent
			TopContent.TopContentClass = "标签分类"
			TopContent.TopContentName = tag.TagName
			TopContent.TopArticle = TagArticles
		}
	}

	//base_right.html
		//获取首页标签

		tag, err := new(cache.Helper).GetAllData(new(cache2.TagCache))
		if err == nil {
			homeContent.Tags = tag.([]*models.Tag)
		} else {
			logger.ZoneLogger.Error("获取Tag错误")
		}

		articleClass, err := new(cache.Helper).GetAllData(new(cache2.ArticleClassCache))
		if err == nil {
			homeContent.ArticleClass = articleClass.([]*models.ArticleClass)
		} else {
			logger.ZoneLogger.Error("获取文章分类错误")
		}

		//获取友情链接
		link, err := new(cache.Helper).GetAllData(new(cache2.LinkCache))
		if err == nil {
			homeContent.Links = link.([]*models.Link)
		} else {
			logger.ZoneLogger.Error("获取友情链接错误")
		}

	homeContent.Articles = Articles
	homeContent.SortType = sortType
	homeContent.ContentType = enum.ContentType(contentType)
	homeContent.TopContent = TopContent
	this.Data["HomeContent"] = homeContent
	this.TplName = "base.html"
}

func (this *ZoneController) Login() {
	if this.IsLogin == true {
		this.Redirect("/", 302)
	}
	this.TplName = "login.html"
}

func (this *ZoneController) Register() {
	this.TplName = "register.html"
}

func (this *ZoneController) Profile() {
	this.MustLogin()
	this.Data["UserInfo"] = this.User
	this.TplName = "profile.html"
}

func (this *ZoneController) AlterPassword() {
	this.MustLogin()
	this.TplName = "alterpassword.html"
}

func (this *ZoneController) AlterData() {
	this.MustLogin()
	this.TplName = "alterdata.html"
}
