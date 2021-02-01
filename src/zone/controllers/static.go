package controllers

import "gozone/library/enum"

type StaticController struct {
	BaseHandler
}

func (this *StaticController) Search() {
	this.Response(enum.DefaultError,"暂时未开放此功能!")
}
