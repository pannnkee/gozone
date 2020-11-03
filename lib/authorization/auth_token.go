package authorization

import (
	"gozone/library/conn"
	"fmt"
	"time"
)

// 添加用户token
// @param token
// @param id 用户ID
// @return err 错误信息
func AddUserToken(token string, id int64) error {
	key := fmt.Sprintf(conn.UserToken, id)
	client := conn.RFStruct.Client(conn.RedisZone)
	_ = client.SetNX(key, token, time.Hour*24)
	return nil
}
