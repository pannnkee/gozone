package models

type HomeContent struct {
	Articles []*Article
	Tags []*Tag
	ArticleClass []*ArticleClass
}
