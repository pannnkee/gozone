package cache

import (
	"Gozone/library/gocache"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/models"
	"fmt"
)

type LinkCache struct{}

func init() {
	linkCache := new(LinkCache)
	err := new(gocache.Helper).PushListCache(linkCache)
	if err != nil {
		panic(err)
	}
}

func (this *LinkCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "link_type", false, "link_type:%v"
}

func (this *LinkCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.Link).Id)
}

func (this *LinkCache) GetAllData() (data interface{}, err error) {
	data, err = new(dao.LinkDao).GetAll()
	return
}

func (this *LinkCache) GetItemData(articleID int64) (data interface{}, err error) {
	return nil, nil
}
