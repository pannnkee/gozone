package cache

import (
	"Gozone/library/cache"
	"Gozone/src/zone/models"
	"fmt"
)

type ArticleClassCache struct {}

func init() {
	articleClassCache := new(ArticleClassCache)
	err := new(cache.Helper).PushListCache(articleClassCache)
	if err != nil {
		panic(err)
	}
}

func (this *ArticleClassCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "articleClass_type", false, "articleClass_type:%v"
}

func (this *ArticleClassCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.ArticleClass).Id)
}

func (this *ArticleClassCache) GetAllData() (data interface{}, err error) {
	data, err = new(models.ArticleClass).GetAllData()
	return
}

func (this *ArticleClassCache) GetItemData(articleID int64) (data interface{}, err error) {
	return nil, nil
}
