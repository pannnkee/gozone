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
	sortType, _ := this.GetInt64("sortType", 1)
	contentType, _ := this.GetInt64("contentType", 0)

	var Articles []models.Article
	var TopContent models.TopContent
	var count int64
	var err error

	if enum.ContentType(contentType) == enum.DefaultType {
		//获取首页文章
		Articles, count, err = models.ArticleInstance.PageListClass(this.Pager.Offset, this.Pager.Limit, sortType, contentType)
		if err != nil {
			this.Response(enum.DefaultError, err.Error())
			return
		}

		for k, v := range Articles {
			article, _ := models.ArticleClassInstance.FindArticleClassName(v.ArticleClass)
			Articles[k].ArticleClassName = article.ClassName
			Articles[k].CreatedTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")
		}
		this.Pager.Count = count
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
