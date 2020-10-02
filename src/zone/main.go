package main

import (
	"Gozone/library/logger"
	_ "Gozone/src/zone/routers"
	"github.com/astaxie/beego"
	"html/template"
)

func main() {
	defer func() {
		logger.ZoneLogger.Sync()
	}()
	beego.AddFuncMap("StarHtml", StarHtml)
	beego.Run()
}

func StarHtml(nums int) (html template.HTML) {
	str := ""
	for i:=0; i<nums; i++ {
		str += "<i class=\"fa fa-star\"></i>"
	}
	html = beego.Str2html(str)
	return
}

