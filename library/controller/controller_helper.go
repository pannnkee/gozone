package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
)

// 解析Controller Json结构体
// @param controller beego控制器
// @param v 解码结构体
// @return err 错误信息
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
