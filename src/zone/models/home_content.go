package models

type HomeContent struct {
	Articles []*Article
	Tags []*Tag
	ArticleClass []*ArticleClass
	PannnkeeZone string `json:"pannnkee_zone"`
}
