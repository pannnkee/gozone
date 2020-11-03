package cache

import (
	"gozone/library/gocache"
	"gozone/src/zone/dao"
	"gozone/src/zone/models"
	"fmt"
)

type EmojiCache struct{}

func init() {
	emojiCache := new(EmojiCache)
	err := new(gocache.Helper).PushListCache(emojiCache)
	if err != nil {
		panic(err)
	}
}

func (this *EmojiCache) CacheConfig() (cacheName string, needItem bool, itemKey string) {
	return "emoji_type", false, "emoji_type:%v"
}

func (this *EmojiCache) PrimaryKey(model interface{}) string {
	return fmt.Sprintf("%v", model.(*models.Emoji).Id)
}

func (this *EmojiCache) GetAllData() (data interface{}, err error) {
	data, err = new(dao.EmojiDao).GetAll()
	return
}

func (this *EmojiCache) GetItemData(articleID int64) (data interface{}, err error) {
	return nil, nil
}
