package conn

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"gozone/library/config"
	"sync"
	"time"
)

var address =  config.GetConfigStr("redis::address","127.0.0.1")
var password = config.GetConfigStr("redis::password","root123")
var port = config.GetConfigStr("redis::port","6379")

func init() {
	RFStruct.Connect(&redis.Options{
		Addr:       fmt.Sprintf("%v:%v", address, port),
		Password:   password,
		DB:         int(RedisZone),
		PoolSize:   500,
		MaxConnAge: time.Minute,
	})
}


type RedisFactory struct {
	ConnectClient map[RedisDatabase]*redis.Client
	ConnectMutex *sync.Mutex
}

func (this *RedisFactory) Init() *RedisFactory {
	this.ConnectClient = make(map[RedisDatabase]*redis.Client)
	this.ConnectMutex = new(sync.Mutex)
	return this
}

// 链接redis客户端
// @param index 库下标
func (this *RedisFactory) Connect(opt *redis.Options) {
	this.ConnectMutex.Lock()
	defer this.ConnectMutex.Unlock()

	client := redis.NewClient(opt)
	if ping := client.Ping().Err(); ping != nil {
		panic(errors.New(fmt.Sprintf("redis connection fail %v", ping.Error())))
	}
	this.ConnectClient[RedisDatabase(opt.DB)] = client
}

// 获取一个redis链接库的客户端
// @param index 库下标文件
func (this *RedisFactory) Client(index RedisDatabase) *redis.Client {

	this.ConnectMutex.Lock()
	defer this.ConnectMutex.Unlock()

	connect, exists := this.ConnectClient[index]
	if !exists {
		panic(errors.New(fmt.Sprintf("undinfed redis connection in %v , please set redis.Client int this.SetConnect", index)))
	}
	return connect
}

