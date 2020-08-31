package config

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"strings"
)

func GetConfigStr(key, defaultStr string) (str string) {
	if beego.AppConfig.String("runmode") == "prod" || strings.ToLower(os.Getenv("BEEGO_RUNMODE")) == "prod" {
		str = beego.AppConfig.String(key)
		if len(str) == 0 {
			logs.Error(fmt.Sprintf("缺少配置[%s]", key))

		}
	} else {
		str = beego.AppConfig.DefaultString(key, defaultStr)
	}
	return
}

func GetConfigInt(key string, def int64) (v int64) {
	if beego.AppConfig.String("runmode") == "prod" || strings.ToLower(os.Getenv("BEEGO_RUNMODE")) == "prod" {
		v, err := beego.AppConfig.Int64(key)
		if v == 0 || err != nil {
			panic(fmt.Sprintf("缺少配置[%s]", key))
		}
		return v
	} else {
		v = beego.AppConfig.DefaultInt64(key, def)
	}
	return
}