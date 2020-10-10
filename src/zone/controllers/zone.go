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

	// 时间排序 热度排序
	sortType, _ := this.GetInt64("sortType", 1)
	// 轮播图模式 top模式(分类 标签)
	contentType, _ := this.GetInt64("contentType", 0)

	var Articles []*models.Article
	var TopContent models.TopContent
	var count int64
	var err error

	homeContent := new(models.HomeContent)
	//默认首页
	if enum.ContentType(contentType) == enum.DefaultType {
		//获取首页文章
		Articles, count, err = models.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
		if err != nil {
			this.Response(enum.DefaultError, err.Error())
			return
		}

		for _, v := range Articles {

			articleClassInterface, _ := new(cache.Helper).GetByItemKey(new(cache2.ArticleClassCache), v.ArticleClass)
			article := articleClassInterface.(*models.ArticleClass)
			commentNums, _ := models.CommentInstance.GetCommentNumsAndHuman(v.Id)

			v.ArticleClassName = article.ClassName
			v.CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
			v.CommentNumber = commentNums
		}
		this.Pager.Count = count
		this.Data["title"] = "PannnKee's Zone"
		this.Data["isHome"] = true
		//获取轮播图

	} else {
	//base_top首页

		// 文章分类
		if enum.ContentType(contentType) < enum.Mysql {

			articleClassInterface, _ := new(cache.Helper).GetByItemKey(new(cache2.ArticleClassCache), contentType)
			class := articleClassInterface.(*models.ArticleClass)
			Articles, _, err = models.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}

			for _, v := range Articles {
				v.CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
			}

			TopContent.ContentNum = class.Nums
			TopContent.ContentText = class.ClassIntroduction
			TopContent.TopContentClass = "文章分类"
			TopContent.TopContentName = class.ClassName
			TopContent.TopArticle = Articles
			this.Data["title"] = "文章分类-PannnKee's Zone"
		} else {
			// 标签分类
			tagInterface, _ := new(cache.Helper).GetByItemKey(new(cache2.TagCache), contentType - 100)
			tag := tagInterface.(*models.Tag)

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
			this.Data["title"] = "标签分类-PannnKee's Zone"
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
		// 获取文章分类
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
	this.Data["title"] = "登录-PannnKee's Zone"
	this.TplName = "login.html"
}

func (this *ZoneController) Register() {
	this.Data["title"] = "注册-PannnKee's Zone"
	this.TplName = "register.html"
}

func (this *ZoneController) Profile() {
	this.MustLogin()
	this.Data["UserInfo"] = this.User
	this.Data["title"] = "个人资料-PannnKee's Zone"
	this.TplName = "profile.html"
}

func (this *ZoneController) AlterPassword() {
	this.MustLogin()
	this.Data["title"] = "修改密码-PannnKee's Zone"
	this.TplName = "alterpassword.html"
}

func (this *ZoneController) AlterData() {
	this.MustLogin()
	this.Data["title"] = "修改资料-PannnKee's Zone"
	this.TplName = "alterdata.html"
}

func (this *ZoneController) TimeLine() {
	data, _ := models.TimeLineInstance.GetAllData()
	for _, v := range data {
		v.UpdateTimeStr = time.Unix(v.UpdateTime, 0).Format("2006-01-02")
	}
	this.Data["Timeline"] = data
	this.Data["title"] = "Timeline-PannnKee's Zone"
	this.TplName = "timeline.html"
}

func (this *ZoneController) About() {
	homeContent := new(models.HomeContent)
	//base_right.html
	//获取首页标签
	tag, err := new(cache.Helper).GetAllData(new(cache2.TagCache))
	if err == nil {
		homeContent.Tags = tag.([]*models.Tag)
	} else {
		logger.ZoneLogger.Error("获取Tag错误")
	}
	// 获取文章分类
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

	aboutData, _ := models.AboutInstance.GetAllData()
	this.Data["AboutData"] = aboutData
	this.Data["HomeContent"] = homeContent
	this.Data["isAbout"] = true
	this.Data["title"] = "关于网站-PannnKee's Zone"
	this.TplName = "about.html"
}

func (this *ZoneController) Archive() {

	articleInterface, _ := new(cache.Helper).GetAllData(new(cache2.ArticleCache))
	article := articleInterface.([]*models.Article)


	//var year models.Year
	//var mouthitem models.MouthItem
	//var articleItem models.ArticleItem
	//var archiveResp models.ArchiveResp
	for k,v := range article {
		fmt.Println(k,v)
	}


	this.Data["isArchive"] = true
	this.Data["title"] = "博客归档-PannnKee's Zone"
	this.TplName = "archive.html"
}