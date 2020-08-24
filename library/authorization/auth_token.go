package authorization

import (
	"Gozone/library/conn"
	"fmt"
	"time"
)

func AddUserToken(token string, id int64) error {
	key := fmt.Sprintf(conn.UserToken, id)
	client := conn.RFStruct.Client(conn.RedisZone)
	_ = client.SetNX(key, token, time.Hour*24)
	return nil
}