package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
)

// 解析Controller Json结构体
func ParseRequestStruct(controller beego.Controller, v interface{}) (err error) {
	req := controller.Ctx.Input.RequestBody
	if len(req) == 0 {
		err = errors.New("RequestBody is nil")
	}
	err = json.Unmarshal(req, v)
	if err != nil {
		err = errors.New(fmt.Sprintf("please checkout params: %v", err.Error()))
	}
	return
}
