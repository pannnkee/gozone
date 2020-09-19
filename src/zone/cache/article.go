package cache

import (
	"Gozone/library/cache"
	"Gozone/src/zone/models"
	"fmt"
)

type ArticleCache struct {}

func init() {
	articleCache := new(ArticleCache)
	err := new(cache.Helper).PushListCache(articleCache)
	if err != nil {
		panic(err)
	}
}

func (this *ArticleCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "article_type", false, "article_type:%v"
}

func (this *ArticleCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.Article).Id)
}

func (this *ArticleCache) GetAllData() (data interface{}, err error) {
	data, err = new(models.Article).GetAllData()
	return
}
