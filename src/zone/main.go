package main

import (
	"Gozone/library/conn"
	_"Gozone/src/zone/routers"
	"github.com/astaxie/beego"
)

func main() {

	conn.GetORMByName("zone")

	beego.Run()
}

