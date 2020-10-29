package cache

import (
	"Gozone/library/cache"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/models"
	"fmt"
)

type TagCache struct {}

func init() {
	tagCache:= new(TagCache)
	err := new(cache.Helper).PushListCache(tagCache)
	if err != nil {
		panic(err)
	}
}

func (this *TagCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "tag_type", false, "tag_type:%v"
}

func (this *TagCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.Tag).Id)
}

func (this *TagCache) GetAllData() (data interface{}, err error) {
	data, err = new(dao.TagDao).GetAll()
	return
}

func (this *TagCache) GetItemData(articleID int64) (data interface{}, err error) {
	return nil, nil
}
