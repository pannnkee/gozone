package cache

import (
	"Gozone/library/cache"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/models"
	"fmt"
)

type UserCache struct {}

func init() {
	userCache := new(UserCache)
	err := new(cache.Helper).PushListCache(userCache)
	if err != nil {
		panic(err)
	}
}

func (this *UserCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "user_type", false, "user_type:%v"
}

func (this *UserCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.User).Id)
}

func (this *UserCache) GetAllData() (data interface{}, err error) {
	data, err = new(dao.UserDao).GetAll()
	return
}

func (this *UserCache) GetItemData(articleID int64) (data interface{}, err error) {
	return nil, nil
}


