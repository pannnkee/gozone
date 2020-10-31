package verifycode

import (
	"Gozone/library/conn"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 添加注册验证码
// @param email 邮箱
// @param code 验证码
// @return err 错误信息
func Add(email string, code string) error {
	key := fmt.Sprintf(conn.UserVerifyCode, email)
	client := conn.RFStruct.Client(conn.RedisZone)
	_ = client.SetNX(key, code, time.Minute*5)
	return nil
}

// 根据邮箱获取验证码
// @param email  邮箱
// @return code 验证码
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

// 根据邮箱删除验证码
// @param email 邮箱
func Del(email string) {
	key := fmt.Sprintf(conn.UserVerifyCode, email)
	client := conn.RFStruct.Client(conn.RedisZone)
	_ = client.Del(key)
}
