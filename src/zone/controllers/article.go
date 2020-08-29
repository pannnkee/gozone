package controllers

import (
	"Gozone/src/zone/models"
	"fmt"
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

