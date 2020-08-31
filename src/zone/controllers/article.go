package controllers

import (
	"Gozone/src/zone/models"
	"fmt"
	"strconv"
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
	this.Response(0,"", datas, this.Pager)
}

func (this *ArticleController) Get() {
	articleIdStr := this.Ctx.Input.Param(":id")
	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)
	if articleId < 1 {
		this.Response(1,"文章参数错误")
		return
	}

	article := new(models.Article)
	err := article.Get(articleId)
	if err != nil {
		this.Response(1,"获取文章错误")
		return
	}

	articleContent := new(models.ArticleContent)
	err = articleContent.Get(articleId)
	if err != nil {
		this.Response(1,"获取文章内容错误")
		return
	}

	data := models.ArticleResp{
		Article:        article,
		ArticleContent: articleContent.Content,
	}
	this.Response(0,"", data)
}

