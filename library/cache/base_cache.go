package cache

import (
	"time"
	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
)

// BaseCache 基本的缓存
type BaseCache struct {

}

// AllCaches 所有缓存
var AllCaches *cache.Cache

//	newCache()
//	创建一个具有给定默认过期时间和清除间隔的新缓存；
//	如果过期时间小于1(或NoExpiration)，则缓存中的项永远不会过期(默认情况下)，必须手动删除。
//	如果清理间隔小于1，则在调用c.DeleteExpired()之前不会从缓存中删除过期的项。
//	[参数]
//	[返回] 给AllCaches赋值
//
func (this *BaseCache) newCache() {
	defaultTime := beego.AppConfig.DefaultInt("go_cache::defaultHour", 48)
	timeInt := time.Duration(defaultTime)
	AllCaches = cache.New(timeInt*time.Hour, 60*time.Second)
}

//	Del(key string)
//	从缓存中删除key的对应项。如果key不在缓存中，则不执行任何操作。
//	[参数]
//	[返回]
//
func (this *BaseCache) Del(key string) {
	if len(key) == 0 {
		return
	}
	if AllCaches == nil {
		this.newCache()
	}
	AllCaches.Delete(key)
}

//	Get(key string) (data interface{}, isExist bool)
//	从缓存中获取key的对应项。返回项或nil，以及是否找到键的bool。
//	[参数]
//	[返回]	isExist：是否找到键的bool；
//
func (this *BaseCache) Get(key string) (data interface{}, isExist bool) {
	if len(key) == 0 {
		return
	}
	if AllCaches == nil {
		this.newCache()
	}
	data, isExist = AllCaches.Get(key)
	return
}

//	Set(key string, data interface{})
//	向缓存中添加key的对应项或替换现有key的对应项。
//	如果持续时间为0 (DefaultExpiration)，则使用缓存的默认过期时间。如果是-1 (NoExpiration)，则该条目永远不会过期。
//	[参数]	key,data:k-v;
//	[返回]
//
func (this *BaseCache) Set(key string, data interface{}) {
	if len(key) == 0 {
		return
	}
	if AllCaches == nil {
		this.newCache()
	}
	AllCaches.Set(key, data, cache.DefaultExpiration)
}

//	DelAll()
//	清空所有设置的缓存
//	[参数]
//	[返回]
//
func (this *BaseCache) DelAll() {
	if AllCaches == nil {
		this.newCache()
	}
	AllCaches.Flush()
}

