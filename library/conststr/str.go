package conststr

import "Gozone/library/config"

var (
	ExpireHour    = config.GetConfigInt("xxtea::expireHour", 168) //token TTL 168小时
	AdminXXTEAKey = config.GetConfigStr("xxtea::jwtkey", "9foklD6p4cv601Dxes")
)
