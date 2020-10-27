package controllers

import (
	"Gozone/library/cache"
	"Gozone/library/enum"
	"Gozone/library/logger"
	cache2 "Gozone/src/zone/cache"
	"Gozone/src/zone/models"
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
			articleClass, _ := models.ArticleClassInstance.Get(v.ArticleClass)
			commentNums, _ := models.CommentInstance.GetCommentNumsAndHuman(v.Id)

			v.ArticleClassName = articleClass.ClassName
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

			articleClass, _ := new(models.ArticleClass).Get(contentType)
			nums, _ := models.ArticleInstance.GetArticleClassNums(articleClass.Id)
			articleClass.Nums = nums

			Articles, _, err = models.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}

			for _, v := range Articles {
				articleClass, _ := models.ArticleClassInstance.Get(v.ArticleClass)
				commentNums, _ := models.CommentInstance.GetCommentNumsAndHuman(v.Id)

				v.ArticleClassName = articleClass.ClassName
				v.CommentNumber = commentNums
				v.CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
			}

			TopContent.ContentNum = articleClass.Nums
			TopContent.ContentText = articleClass.ClassIntroduction
			TopContent.TopContentClass = "文章分类"
			TopContent.TopContentName = articleClass.ClassName
			TopContent.TopArticle = Articles
			this.Data["title"] = "文章分类-PannnKee's Zone"
		} else {
			// 标签分类

			var Tag[]int64
			tag, _ := models.TagInstance.Get(contentType - 100)
			articles, _ := models.ArticleTagInstance.FindArticles(tag.Id)
			for _, v := range articles {
				Tag = append(Tag, v.ArticleId)
			}

			TagArticles, err := models.ArticleInstance.FindArticles(Tag)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}
			for _, v := range TagArticles {
				articleClass, _ := models.ArticleClassInstance.Get(v.ArticleClass)
				commentNums, _ := models.CommentInstance.GetCommentNumsAndHuman(v.Id)

				v.ArticleClassName = articleClass.ClassName
				v.CommentNumber = commentNums
				v.CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
			}
			TopContent.ContentNum = int64(len(articles))
			TopContent.ContentText = tag.TagContent
			TopContent.TopContentClass = "标签分类"
			TopContent.TopContentName = tag.TagName
			TopContent.TopArticle = TagArticles
			this.Data["title"] = "标签分类-PannnKee's Zone"
		}
	}

	//base_right.html
		//获取首页标签

		tags, err := models.TagInstance.GetAllData()
		if err == nil {
			homeContent.Tags = tags
		} else {
			logger.ZoneLogger.Error("获取Tag错误")
		}
		// 获取文章分类
		articleClass, _ := new(models.ArticleClass).GetAllData()
		if err == nil {
			for _, v := range articleClass {
				nums, _ := models.ArticleInstance.GetArticleClassNums(v.Id)
				v.Nums = nums
			}
			homeContent.ArticleClass = articleClass
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
	this.Data["Description"] = "PannnKee's Zone是一个Beego搭建的博客，分享学习心得、经验总结。主要包括Golang、Docker、Kubernetes、ServerMesh等。"
	this.Data["Keywords"] = "Golang编程，Docker入门，Kubernetes搭建，Golang Web开发，个人博客"
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
	tags, err := models.TagInstance.GetAllData()
	if err == nil {
		homeContent.Tags = tags
	} else {
		logger.ZoneLogger.Error("获取Tag错误")
	}
	// 获取文章分类
	articleClass, _ := new(models.ArticleClass).GetAllData()
	if err == nil {
		for _, v := range articleClass {
			nums, _ := models.ArticleInstance.GetArticleClassNums(v.Id)
			v.Nums = nums
		}
		homeContent.ArticleClass = articleClass
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

	homeContent := new(models.HomeContent)
	//base_right.html
	//获取首页标签
	tags, err := models.TagInstance.GetAllData()
	if err == nil {
		homeContent.Tags = tags
	} else {
		logger.ZoneLogger.Error("获取Tag错误")
	}
	// 获取文章分类
	articleClass, _ := new(models.ArticleClass).GetAllData()
	if err == nil {
		for _, v := range articleClass {
			nums, _ := models.ArticleInstance.GetArticleClassNums(v.Id)
			v.Nums = nums
		}
		homeContent.ArticleClass = articleClass
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

	this.Data["isArchive"] = true
	this.Data["title"] = "博客归档-PannnKee's Zone"
	this.Data["HomeContent"] = homeContent
	this.TplName = "archive.html"
}