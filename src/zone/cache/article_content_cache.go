package cache

import (
	"Gozone/library/cache"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/models"
	"fmt"
)

type ArticleContentCache struct {}

func init() {
	articleContentCache := new(ArticleContentCache)
	err := new(cache.Helper).PushListCache(articleContentCache)
	if err != nil {
		panic(err)
	}
}

func (this *ArticleContentCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "articleContent_type", false, "articleContent_type:%v"
}

func (this *ArticleContentCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.ArticleContent).Id)
}

func (this *ArticleContentCache) GetAllData() (data interface{}, err error) {
	data, err = new(dao.ArticleContentDao).GetAll()
	return
}

func (this *ArticleContentCache) GetItemData(articleID int64) (data interface{}, err error) {
	return nil, nil
}
