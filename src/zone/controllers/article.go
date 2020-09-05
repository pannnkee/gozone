package controllers

import (
	"Gozone/library/util"
	"Gozone/src/zone/models"
	"fmt"
	"html"
	"strconv"
	"time"
)

type ArticleController struct {
	BaseHandler
}

func (this *ArticleController) PageList() {

	datas, count, err := models.ArticleInstance.PageList(this.Pager.Offset, this.Pager.Limit)
	if err != nil {
		this.Response(1, fmt.Sprintf("查询错误:%v", err))
	}

	this.Pager.Count = count
	this.Response(0, "", datas, this.Pager)
}

func (this *ArticleController) Get() {
	articleIdStr := this.Ctx.Input.Param(":id")
	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)
	if articleId < 1 {
		this.Response(1, "文章参数错误")
		return
	}

	article := new(models.Article)
	err := article.Get(articleId)
	if err != nil {
		this.Response(1, "获取文章错误")
		return
	}

	article.CreatedTimeStr = time.Unix(article.CreateTime,0).Format("2006-01-02")
	article.UpdateTimeStr = time.Unix(article.UpdateTime,0).Format("2006-01-02")

	articleContent := new(models.ArticleContent)
	err = articleContent.Get(articleId)
	if err != nil {
		this.Response(1, "获取文章内容错误")
		return
	}

	articleTag := new(models.ArticleTag)
	signs, err := articleTag.FindTags(articleId)
	if err != nil {
		this.Response(1, "获取文章标签错误")
		return
	}

	ArticleClass, err := new(models.ArticleClass).FindArticleName(article.ArticleClass)
	if err != nil {
		this.Response(1,"获取文章类别名称错误")
		return
	}

	tagNames := []string{}
	for _, v := range signs {
		tag := new(models.Tag)
		tagName, _ := tag.GetTagName(v.TagId)
		tagNames = append(tagNames, tagName)
	}
	
	data := models.ArticleListResp{
		Article:        *article,
		ArticleContent: html.UnescapeString(articleContent.Content),
		ArticleTags:    tagNames,
		ArticleClassName: ArticleClass.ClassName,
	}
	jsonMap, err := util.Struct2JsonMap(data)
	if err != nil {
		this.Response(1,fmt.Sprintf("序列化错误:%v", err.Error()))
		return
	}
	this.Data["articleResp"] = jsonMap
	this.TplName = "article.html"
}
