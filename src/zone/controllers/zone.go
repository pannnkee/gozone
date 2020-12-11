package controllers

import (
	"fmt"
	"gozone/library/enum"
	"gozone/library/gocache"
	"gozone/library/logger"
	"gozone/library/util"
	cache2 "gozone/src/zone/cache"
	"gozone/src/zone/dao"
	"gozone/src/zone/models"
	"strconv"
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
		Articles, count, err = dao.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
		if err != nil {
			this.Response(enum.DefaultError, err.Error())
			return
		}

		for _, v := range Articles {
			articleClass, _ := dao.ArticleClassInstance.Get(v.ArticleClass)
			commentNums, _ := dao.CommentInstance.GetCommentNumsAndHuman(v.Id)

			v.ArticleClassName = articleClass.ClassName
			v.CreatedTimeStr = time.Unix(v.CreateTime, 0).Format("2006-01-02")
			v.CommentNumber = commentNums
		}
		this.Pager.Count = count
		this.Data["title"] = "PannnKee's Zone"
		this.Data["isHome"] = true

		//获取轮播图
		articleCarousel, _ := dao.ArticleInstance.GetCarouselArticle()
		for _, v := range articleCarousel {
			homeContent.CarouselArticle = append(homeContent.CarouselArticle, v)
		}

	} else {
		//base_top首页

		// 文章分类
		if enum.ContentType(contentType) < enum.Mysql {

			articleClass, _ := dao.ArticleClassInstance.Get(contentType)
			nums, _ := dao.ArticleInstance.GetArticleClassNums(articleClass.Id)
			articleClass.Nums = nums

			Articles, _, err = dao.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}

			for _, v := range Articles {
				articleClass, _ := dao.ArticleClassInstance.Get(v.ArticleClass)
				commentNums, _ := dao.CommentInstance.GetCommentNumsAndHuman(v.Id)

				v.ArticleClassName = articleClass.ClassName
				v.CommentNumber = commentNums
				v.CreatedTimeStr = time.Unix(v.CreateTime, 0).Format("2006-01-02")
			}

			TopContent.ContentNum = articleClass.Nums
			TopContent.ContentText = articleClass.ClassIntroduction
			TopContent.TopContentClass = "文章分类"
			TopContent.TopContentName = articleClass.ClassName
			TopContent.TopArticle = Articles
			this.Data["title"] = "文章分类-PannnKee's Zone"
		} else {
			// 标签分类

			var Tag []int64
			tag, _ := dao.TagInstance.Get(contentType - 100)
			articles, _ := dao.ArticleTagInstance.FindArticles(tag.Id)
			for _, v := range articles {
				Tag = append(Tag, v.ArticleId)
			}

			TagArticles, err := dao.ArticleInstance.FindArticles(Tag)
			if err != nil {
				this.Response(enum.DefaultError, err.Error())
				return
			}
			for _, v := range TagArticles {
				articleClass, _ := dao.ArticleClassInstance.Get(v.ArticleClass)
				commentNums, _ := dao.CommentInstance.GetCommentNumsAndHuman(v.Id)

				v.ArticleClassName = articleClass.ClassName
				v.CommentNumber = commentNums
				v.CreatedTimeStr = time.Unix(v.CreateTime, 0).Format("2006-01-02")
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
	tags, _ := dao.TagInstance.GetAll()
	homeContent.Tags = tags

	// 获取文章分类
	articleClass, _ := dao.ArticleClassInstance.GetAll()
	for _, v := range articleClass {
		nums, _ := dao.ArticleInstance.GetArticleClassNums(v.Id)
		v.Nums = nums
	}
	homeContent.ArticleClass = articleClass

	//获取友情链接
	link, _ := new(gocache.Helper).GetAllData(new(cache2.LinkCache))
	homeContent.Links = link.([]*models.Link)

	nums, _ := new(dao.ArticleDao).GetArticleNums()
	pageInstance := util.HtmlPage(this.Pager.Page, nums, this.Pager.Limit, sortType)
	// 默认首页才有分页
	if enum.ContentType(contentType) == enum.DefaultType {
		pageInstance.IsShow = true
	}

	homeContent.Articles = Articles
	homeContent.SortType = sortType
	homeContent.ContentType = enum.ContentType(contentType)
	homeContent.TopContent = TopContent
	this.Data["HomeContent"] = homeContent
	this.Data["PageInstance"] = pageInstance
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
	data, _ := dao.TimeLineInstance.GetAll()
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
	tags, err := dao.TagInstance.GetAll()
	if err == nil {
		homeContent.Tags = tags
	} else {
		logger.ZoneLogger.Error("获取Tag错误")
	}
	// 获取文章分类
	articleClass, _ := dao.ArticleClassInstance.GetAll()
	if err == nil {
		for _, v := range articleClass {
			nums, _ := dao.ArticleInstance.GetArticleClassNums(v.Id)
			v.Nums = nums
		}
		homeContent.ArticleClass = articleClass
	} else {
		logger.ZoneLogger.Error("获取文章分类错误")
	}

	//获取友情链接
	link, err := new(gocache.Helper).GetAllData(new(cache2.LinkCache))
	if err == nil {
		homeContent.Links = link.([]*models.Link)
	} else {
		logger.ZoneLogger.Error("获取友情链接错误")
	}

	aboutData, _ := dao.AboutInstance.GetAllData()
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
	tags, err := dao.TagInstance.GetAll()
	if err == nil {
		homeContent.Tags = tags
	} else {
		logger.ZoneLogger.Error("获取Tag错误")
	}
	// 获取文章分类
	articleClass, _ := dao.ArticleClassInstance.GetAll()
	if err == nil {
		for _, v := range articleClass {
			nums, _ := dao.ArticleInstance.GetArticleClassNums(v.Id)
			v.Nums = nums
		}
		homeContent.ArticleClass = articleClass
	} else {
		logger.ZoneLogger.Error("获取文章分类错误")
	}

	//获取友情链接
	link, err := new(gocache.Helper).GetAllData(new(cache2.LinkCache))
	if err == nil {
		homeContent.Links = link.([]*models.Link)
	} else {
		logger.ZoneLogger.Error("获取友情链接错误")
	}


	//拼接归档html
	articleYearMap := make(map[int]map[time.Month][]*models.Article, 0)
	allArticle, _ := dao.ArticleInstance.GetAllDataByCreateTime()

	for _, v := range allArticle {
		tempMouth := articleYearMap[time.Unix(v.CreateTime, 0).Year()]
		if tempMouth == nil {
			tempMouth = make(map[time.Month][]*models.Article, 0)
		}
		tempMouth[time.Unix(v.CreateTime, 0).Month()] = append(tempMouth[time.Unix(v.CreateTime, 0).Month()], v)
		articleYearMap[time.Unix(v.CreateTime,0).Year()] = tempMouth
	}

	html := ""
	for k,v := range articleYearMap {
		html += "<li>" + strconv.Itoa(k) + "年"
		html += "<u1 class= \"pl-4\">"
			for index, value := range v {
				html += "<li>" + strconv.Itoa(int(index)) + "共 " + strconv.Itoa(len(value)) + " 篇"
					for _, a := range value {
						html += "<ul class=\"pl-4\">"
							html += fmt.Sprintf("<li class=\"text-info\">%v-%v&nbsp;&nbsp;<a href=\"/article/%v\">%v</a></li>",
								int(time.Unix(a.CreateTime,0).Month()),time.Unix(a.CreateTime,0).Day(), a.Id, a.ArticleTitle)
						html += "</ul>"
					}
				html += "</u1>"
			}
		html += "</u1>"
		html += "</li>"
	}

	this.Data["isArchive"] = true
	this.Data["title"] = "博客归档-PannnKee's Zone"
	this.Data["HomeContent"] = homeContent
	this.Data["archiveHTML"] = html
	this.TplName = "archive.html"
}
