package main

import (
	"github.com/astaxie/beego"
	"gozone/library/logger"
	_ "gozone/src/zone/routers"
	"html/template"
)

func main() {
	defer func() {
		logger.ZoneLogger.Sync()
	}()
	beego.AddFuncMap("StarHtml", StarHtml)
	beego.Run()
}

// Timeline中事件星级数量
// @pram nums 星星数量
// @return html 返回的html
func StarHtml(nums int) (html template.HTML) {
	str := ""
	for i := 0; i < nums; i++ {
		str += "<i class=\"fa fa-star\"></i>"
	}
	html = beego.Str2html(str)
	return
}
