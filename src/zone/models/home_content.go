package models

import "Gozone/library/enum"

type HomeContent struct {
	Tags []Tag
	Links []Link
	Articles []Article
	SortType int64
	ContentType enum.ContentType
	ArticleClass []ArticleClass
	PannnkeeZone string `json:"pannnkee_zone"`
}
