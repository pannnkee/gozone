package util

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"strings"
)

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

//func MarkDown2Html() string {
//	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
//	parser := parser.NewWithExtensions(extensions)
//	md := []byte("由于被疫情打乱更博节奏，已经好久没有更新博客了，今天来一发笔记，记录一下 linux 三剑客的常用操作。所谓的三剑客，其实就是 `grep` `awk` `sed` 三个命令，在 linux 操作中，特别是关于文本的查找和修改，这三条命令可谓是各显神通。\n\n三个命令的一般性作用：\n\n*   grep：基于正则表达式查找满足条件的行\n*   awk：根据定位到的数据行处理其中的分段，也就是切片\n*   sed：根据定位到的数据行修改数据\n\n## sed 命令常用操作\n\nsed 命令的匹配是以行为单位的，这个概念很重要。\n\n直接上测试文件 file.conf，内容如下：\n\n```\nfunction test() {\n\tconsole.log(\"Hello world!\");\n}\n```\n\n### 简单替换（只替换每个匹配行的首次匹配项）\n\n    sed ‘s/book/good/‘ file.conf\n\n结果如下：\n\n```\nfunc (this *ArticleCache) GetAllData() (data interface{}, err error) {\n\tdata, err = new(models.Article).GetAllData()\n\treturn\n}\n```\n\n可以看到只将每个匹配到的行的第一次匹配到的 book 并替换成了 good\n\n### 精确替换\n\n上面的替换虽然可以做到将book替换成新的内容，但是有时候我们可能需要对某个单词进行替换，而不能替换包含这个单词的单词，就比如上面的内容，假如现在要将 bookbook 替换成 newbook ，先看看使用上面的替换方式结果如何\n\n    sed ‘s/bookbook/newword/‘ file.conf\n\n结果如下\n\n```shell\nnewwordbookbook\nname=jack\nip=127.0.0.1\nurl=https://abc.xxxxx.com\nxyz showowof3442 lmn\nnewword\n```")
//	html := markdown.ToHTML(md, parser, nil)
//	return string(html)
//}
