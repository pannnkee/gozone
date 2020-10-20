package controllers

import (
	"Gozone/library/cache"
	"Gozone/library/config"
	"Gozone/library/controller"
	"Gozone/library/enum"
	"Gozone/library/logger"
	"Gozone/library/util"
	cache2 "Gozone/src/zone/cache"
	"Gozone/src/zone/model_view"
	"Gozone/src/zone/models"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ArticleController struct {
	BaseHandler
}

var EmojiMap map[string]*models.Emoji

func init() {
	EmojiMap = make(map[string]*models.Emoji)
	data, err := models.EmojiInstance.GetAllData()
	if err != nil {
		logger.ZoneLogger.Error("初始化EmojiMap错误:", err.Error())
	} else {
		for _, v := range data {
			EmojiMap[v.DataEmoji] = v
		}
	}
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

	//文章观看次数+1
	_ = new(models.Article).UpdateViews(articleId)
	//_ = new(cache.Helper).UpDataItem(new(cache2.ArticleCache), articleId)

	wg := new(sync.WaitGroup)
	data := models.ArticleListResp{}

	wg.Add(5)

	go func() {
		now := time.Now()
		//文章详情
		defer wg.Done()
		article := new(models.Article)
		err := article.Get(articleId)
		if err == nil {
			article.CreatedTimeStr = time.Unix(article.CreateTime, 0).Format("2006-01-02 15:04:05")
			article.UpdateTimeStr = time.Unix(article.UpdateTime, 0).Format("2006-01-02 15:04:05")
			data.Article = *article
		} else {
			logger.ZoneLogger.Error("获取文章详情错误")
		}

		// 获取文章分类
		articleClass, _ := new(models.ArticleClass).Get(articleId)
		data.ArticleClassName = articleClass.ClassName

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

		articleContent := new(models.ArticleContent)
		err := articleContent.Get(articleId)
		if err == nil {
			down2Html := util.MarkDown2Html(articleContent.Content)
			data.ArticleContent = down2Html

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
			v.CreateTimeStr = time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:05")
			v.Content = Emoji2Html(v.Content)
			v.Content = util.MarkDown2Html(v.Content)

			user, err := models.UserInstance.Get(v.UserID)
			if err == nil {
				v.UserAvatar = user.Avatar
				if v.UserAvatar == "" {
					v.UserAvatar = "/static/img/user_avatar/default_avatar.png"
				}
			}

			secondComment, err := models.CommentInstance.GetSecondComment(articleId, v.ID)
			if err == nil {
				for _, value := range secondComment {
					value.Content = Emoji2Html(value.Content)
					value.Content = util.MarkDown2Html(value.Content)
					value.CreateTimeStr = time.Unix(value.CreateTime, 0).Format("2006-01-02 15:04:05")

					user, err := models.UserInstance.Get(value.UserID)
					if err == nil {
						value.UserAvatar = user.Avatar
						if value.UserAvatar == "" {
							value.UserAvatar = "/static/img/user_avatar/default_avatar.png"
						}
					}
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
	this.Data["title"] = fmt.Sprintf("%v-PannnKee's Zone", data.Article.ArticleTitle)
	this.Data["articleURL"] = config.GetConfigStr("zone:site", "http://127.0.0.1") + this.RequestURL
	this.Data["articleResp"] = jsonMap
	this.TplName = "article.html"
}

func (this *ArticleController) Comment() {
	if !this.IsLogin {
		this.Response(enum.DefaultError, "请登录账号后再评论")
	}

	//TODO 关键词屏蔽
	commentWeb := new(model_view.CommentWeb)
	err := controller.ParseRequestStruct(this.Controller, &commentWeb)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
	}

	now := time.Now().Unix()
	//盖楼
	if commentWeb.RepID == 0 {
		comment := models.Comment{
			IP:            this.GetIP(),
			UserID:        this.User.Id,
			UserName:      this.User.UserName,
			ArticleID:     commentWeb.ArticleId,
			CommentLevel:  1,
			Content:       commentWeb.Content,
			Status:        1,
			CreateTime:    now,
			CreateTimeStr: time.Unix(now, 0).Format("2006-01-02"),
		}
		err := comment.AddComment()
		if err != nil {
			this.Response(enum.DefaultError, err.Error())
		}
	} else {
		//回复评论
		comment := models.Comment{
			UserID:          this.User.Id,
			UserName:        this.User.UserName,
			ArticleID:       commentWeb.ArticleId,
			ParentCommentID: commentWeb.RepID,
			ReplyCommentID:  commentWeb.ReplyFatherID,
			CommentLevel:    2,
			Content:         commentWeb.Content,
			Status:          1,
			CreateTime:      now,
			CreateTimeStr:   time.Unix(now, 0).Format("2006-01-02"),
		}
		if commentWeb.ReplyFatherID != 0 {
			//说明是二级评论的回复
			user, err := models.UserInstance.Get(commentWeb.RepUserID)
			if err == nil {
				comment.ReplyCommentUserID = user.Id
				comment.ReplyCommentUserName = user.UserName
			}
		}
		err := comment.AddComment()
		if err != nil {
			this.Response(enum.DefaultError, err.Error())
		}
	}
	this.Response(enum.DefaultSuccess, "1")
}

func Emoji2Html(comment string) (html string) {
	for _, v := range EmojiMap {
		if strings.Contains(comment, v.DataEmoji) {
			if emoji, ok := EmojiMap[v.DataEmoji]; ok {
				replaceHTML := fmt.Sprintf("<img class=\"comment-emoji-img\" src=\"%v\" title=\"%v\" alt=\"%v\" data-emoji=\"%v\">",
					emoji.Src, emoji.Title, emoji.Alt, emoji.DataEmoji)
				comment = strings.Replace(comment, v.DataEmoji, replaceHTML, -1)
			}
		}
	}
	return comment
}
