package models

type HomeContent struct {
	Articles []*Article
	Tags []*Tag
	ArticleClass []*ArticleClass
	Links []*Link
	PannnkeeZone string `json:"pannnkee_zone"`
	SortType int64
}
