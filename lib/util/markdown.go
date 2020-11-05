package util

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"strings"
)

// 将markDown内容转换成html 并且将code高亮
func MarkDown2Html(Content string) string {

	// 替换数据库回车 \r\n   --> \n
 	Content = strings.Replace(Content, "\r", "", -1)

	md := []byte(Content)
	markdown := blackfriday.Run(md)


	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	  对document进程查询，选择器和css的语法一样
	  第一个参数：i是查询到的第几个元素
	  第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
	})

	htmlString, _ := doc.Html()
	return htmlString
}