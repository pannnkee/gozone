package verifycode

import (
	"Gozone/library/conn"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Add(email string, code string) error {
	key := fmt.Sprintf(conn.UserVerifyCode, email)
	client := conn.RFStruct.Client(conn.RedisZone)
	_ = client.SetNX(key, code, time.Minute*5)
	return nil
}

func Get(email string) (code string) {
	key := fmt.Sprintf(conn.UserVerifyCode, email)
	client := conn.RFStruct.Client(conn.RedisZone)
	cmd := client.Get(key)
	if cmd.Err() == redis.Nil {
		code = ""
	} else {
		code = cmd.Val()
	}
	return
}

func Del(email string)  {
	key := fmt.Sprintf(conn.UserVerifyCode, email)
	client := conn.RFStruct.Client(conn.RedisZone)
	_ = client.Del(key)
}