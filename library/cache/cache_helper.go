package cache

import (
	"Gozone/library/logger"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"reflect"
	"strconv"
)

type Helper struct {
	BaseCache
}

var itemCacheMap map[string]*cache.Cache = make(map[string]*cache.Cache, 0)

//DB-MODEL普通类实现该接口
type Model interface {
	//返回主键名称
	PrimaryKey (model interface{}) string
	//返回cache配置
	CacheConfig() (cacheName string, needItem bool, itemKey string)
	//查询全数据
	GetAllData() (datas interface{}, err error)
	//更新缓存
	//UpDataItem() (id int64, err error)
	//获取指定数据
	GetItemData(id int64) (data interface{}, err error)
}

func (this *Helper) PushListCache(dataModel Model) (err error) {
	cacheName, needItem, itemKey := dataModel.CacheConfig()
	// 查询全部数据
	data, err := dataModel.GetAllData()
	// 存储全部列表
	this.Set(cacheName, data)
	logger.ZoneLogger.Info(fmt.Sprintf("%v 设置成功", cacheName))

	if needItem {
		itemCache, ok := itemCacheMap[cacheName]
		if ok {
			itemCache.Flush()
		}
		itemCache = cache.New(-1, -1)
		if !ok {
			itemCacheMap[cacheName] = itemCache
		}
		t := reflect.TypeOf(data)
		sliceLen := 0
		if t.Kind() == reflect.Slice {
			v := reflect.ValueOf(data)
			sliceLen = v.Len()
			for i:=0; i<v.Len(); i++ {
				itemData := v.Index(i).Interface()
				itemCache.Set(fmt.Sprintf(itemKey, dataModel.PrimaryKey(itemData)), itemData, -1)
			}
		} else {
			return errors.New(fmt.Sprintf("%v GetAllData方法返回不是slice数据,%v", cacheName, t.Elem().Kind().String()))
		}
		fmt.Println(fmt.Sprintf("%v列表单个缓存加入完毕，共加入%v条", cacheName, sliceLen))
	}
	return
}

func (this *Helper) GetAllData(dataModel Model) (datas interface{}, err error) {
	cacheName, _, _ := dataModel.CacheConfig()
	datas, isExist := this.Get(cacheName)
	if !isExist {
		datas, err = dataModel.GetAllData()
		pErr := this.PushListCache(dataModel)
		if pErr != nil {
			logger.ZoneLogger.Error(pErr)
		}
	}
	return
}

func (s *Helper) GetByItemKey(dataModel Model, itemId int64) (data interface{}, err error) {
	//判断是否需要存储单个元素
	cacheName, needItem, itemKey := dataModel.CacheConfig()
	if !needItem {
		datas, err := s.GetAllData(dataModel)
		if err != nil {
			return nil, err
		}
		t := reflect.TypeOf(datas)
		if t.Kind() == reflect.Slice {
			v := reflect.ValueOf(datas)
			idstr := strconv.Itoa(int(itemId))
			for i := 0; i < v.Len(); i++ {
				itemData := v.Index(i).Interface()
				if dataModel.PrimaryKey(itemData) == idstr {
					return itemData, nil
				}
			}
		} else {
			return nil, errors.New(fmt.Sprintf(cacheName) + " 缓存储存非切片数据")
		}
	} else {
		itemCache, ok := itemCacheMap[cacheName]
		if ok {
			data, isexist := itemCache.Get(fmt.Sprintf(itemKey, itemId))
			if !isexist {
				return nil, errors.New(fmt.Sprintf(itemKey, itemId) + " 缓存不存在")
			} else {
				return data, nil
			}

		} else {
			return nil, errors.New(fmt.Sprintf(itemKey, itemId) + " 缓存不存在")
		}
	}
	return nil, errors.New("缓存不存在")
}

func (s *Helper) UpDataItem(dataModel Model, itemId int64) (err error) {
	cacheName, _, itemKey := dataModel.CacheConfig()

	data, err := dataModel.GetItemData(itemId)
	if err != nil {
		return err
	}
	itemCache, _ := itemCacheMap[cacheName]
	_ = itemCache.Replace(fmt.Sprintf(itemKey, dataModel.PrimaryKey(data), -1), data, -1)
	return nil
}



