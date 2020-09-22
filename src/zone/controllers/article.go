package controllers

import (
	"Gozone/library/cache"
	"Gozone/library/controller"
	"Gozone/library/enum"
	"Gozone/library/logger"
	"Gozone/library/util"
	cache2 "Gozone/src/zone/cache"
	"Gozone/src/zone/models"
	"fmt"
	"html"
	"strconv"
	"sync"
	"time"
)

type ArticleController struct {
	BaseHandler
}

func (this *ArticleController) PageList() {

	typeId, _ := this.GetInt64("type", 0)
	data, count, err := models.ArticleInstance.PageList(this.Pager.Offset, this.Pager.Limit, typeId)
	if err != nil {
		this.Response(1, fmt.Sprintf("查询错误:%v", err))
	}

	this.Pager.Count = count
	this.Response(0, "", data, this.Pager)
}

func (this *ArticleController) Get() {
	articleIdStr := this.Ctx.Input.Param(":id")
	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)
	if articleId < 1 {
		this.Response(1, "文章参数错误")
		return
	}

	wg := new(sync.WaitGroup)
	data := models.ArticleListResp{}

	wg.Add(5)

	go func() {
		now := time.Now()
		//文章详情
		defer wg.Done()

		articleInterface, err := new(cache.Helper).GetByItemKey(new(cache2.ArticleCache), articleId)
		article := articleInterface.(*models.Article)
		if err == nil {
			article.CreatedTimeStr = time.Unix(article.CreateTime, 0).Format("2006-01-02")
			article.UpdateTimeStr = time.Unix(article.UpdateTime, 0).Format("2006-01-02")
			data.Article = *article

			articleClassInterface, _ := new(cache.Helper).GetByItemKey(new(cache2.ArticleClassCache), article.ArticleClass)
			ArticleClass := articleClassInterface.(*models.ArticleClass)
			data.ArticleClassName = ArticleClass.ClassName
		} else {
			logger.ZoneLogger.Error("获取文章详情错误")
		}

		// 评论数 参与人数
		commentNums, Humans := models.CommentInstance.GetCommentNumsAndHuman(articleId)
		data.ArticleContentNums = commentNums
		data.ArticleHumans = Humans
		fmt.Println("文章详情:", time.Since(now))
	}()

	go func() {
		now := time.Now()
		//文章内容
		defer wg.Done()
		article, err := new(cache.Helper).GetByItemKey(new(cache2.ArticleContentCache), articleId)
		if err == nil {
			data.ArticleContent = html.UnescapeString(article.(*models.ArticleContent).Content)
		} else {
			logger.ZoneLogger.Error("获取文章内容错误")
		}
		fmt.Println("文章内容:", time.Since(now))
	}()

	go func() {
		now := time.Now()
		//文章标签
		defer wg.Done()

		articleTag := new(models.ArticleTag)
		signs, err := articleTag.FindTags(articleId)
		if err == nil {
			var tagIds []int64
			for _, v := range signs {
				tagIds = append(tagIds, v.TagId)
			}
			tag := new(models.Tag)
			tags, err := tag.GetTags(tagIds)
			if err == nil {
				data.ArticleTags = tags
			} else {
				logger.ZoneLogger.Error("获取文章标签错误")
			}
		} else {
			logger.ZoneLogger.Error("获取文章标签错误")
		}
		fmt.Println("文章标签:", time.Since(now))
	}()

	go func() {
		// 获取Emoji
		now := time.Now()
		defer wg.Done()
		allData, err := new(cache.Helper).GetAllData(new(cache2.EmojiCache))
		if err == nil {
			tempEmoji := new([]*models.Emoji)
			i := 0
			for _, v := range allData.([]*models.Emoji) {
				*tempEmoji = append(*tempEmoji, v)
				i++

				if i == 8 {
					data.Emoji = append(data.Emoji, *tempEmoji)
					*tempEmoji = nil
					i = 0
				}
			}
		} else {
			logger.ZoneLogger.Error("获取Emoji错误")
		}
		fmt.Println("emoji:", time.Since(now))
	}()



	go func() {
		// 获取文章评论
		now := time.Now()
		defer wg.Done()
		comment, err := models.CommentInstance.GetFirstComment(articleId)
		if err != nil {
			this.Response(1, fmt.Sprintf("获取文章评论错误:%v", err.Error()))
			return
		}

		for k, v := range comment {
			v.Floor = int64(len(comment) - k)
			v.CreateTimeStr = time.Unix(v.CreateTime,0).Format("2006-01-02")

			secondComment, err := models.CommentInstance.GetSecondComment(articleId, v.ID)
			if err == nil {
				for _, value := range secondComment {
					value.CreateTimeStr = time.Unix(value.CreateTime,0).Format("2006-01-02")
				}
				v.SecondComment = secondComment
			}
		}

		data.Comment = comment
		fmt.Println("文章评论:", time.Since(now))
	}()

	wg.Wait()

	jsonMap, err := util.Struct2JsonMap(data)
	if err != nil {
		this.Response(1, fmt.Sprintf("序列化错误:%v", err.Error()))
		return
	}
	this.Data["articleResp"] = jsonMap
	this.TplName = "article.html"
}

func (this *ArticleController) Comment() {
	type A struct {
		RepID     int    `json:"rep_id"`
		Content   string `json:"content"`
		ArticleId int64  `json:"article_id"`
	}
	var a A
	err := controller.ParseRequestStruct(this.Controller, &a)
	fmt.Print(err)

	this.Response(enum.DefaultSuccess, "1")
}
