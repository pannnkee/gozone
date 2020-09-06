package main

import (
	"Gozone/library/logger"
	_ "Gozone/src/zone/routers"
	"github.com/astaxie/beego"
)

func main() {

	defer func() {
		logger.ZoneLogger.Sync()
	}()


	beego.Run()
}

