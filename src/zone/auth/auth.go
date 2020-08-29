package auth

import "github.com/astaxie/beego"

const (
	ZoneToken = "zone_token"
)

var (
	XXTEKEY = beego.AppConfig.String("xxtea::key")
)
