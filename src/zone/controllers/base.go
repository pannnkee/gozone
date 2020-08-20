package controllers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"time"
)

//BaseHandler BaseHandler
type BaseHandler struct {
	beego.Controller
	ClientIP       string
	UserAgent      string
	Method         string
	Host           string
	Body           string
	RequestURL     string
	ControllerName string
	ActionName     string
	Pager          *Pager
}

// DataResponse 返回数据模型
type DataResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Stime string `json:"stime"`
	Body  *Body  `json:"body"`
}

// Body 具体数据模型
type Body struct {
	Pager *Pager      `json:"pager"`
	Data  interface{} `json:"data"`
}

//Pager Pager
type Pager struct {
	Page      int64 `json:"page"`      //页数
	Limit     int64 `json:"limit"`     //大小
	Offset    int64 `json:"offset"`    //偏移量
	Count     int64 `json:"count"`     //总数
	PageCount int64 `json:"pagecount"` //当前页数量
}

func (this *BaseHandler) Response(code int, msg string, args ...interface{}) {
	var (
		err error
		pager *Pager
		data interface{}
	)

	for _, v := range args {
		switch vv := v.(type) {
		case error:
			err = error(vv)
		case *Pager:
			pager = vv
		default:
			data = vv
		}
	}

	//处理error 可记录到日志
	if err != nil {
		log.Errorf("%v | ERROR: %v", this.RequestURL, err)
	}

	if code == 0 && len(msg) == 0 {
		msg = "success"
	}

	resp := &DataResponse{
		Code: code,
		Msg: msg,
		Stime: time.Now().Format("2006-01-02 15:04:05"),
	}

	body := new(Body)
	if pager != nil {
		pager.PageCount = this.GetPageCount(pager.Count, pager.Limit)
		body.Pager = pager
	} else {
		body.Pager = &Pager{}
	}
	body.Data = data
	resp.Body = body

	this.Data["json"] = *resp
	this.ServeJSON()
	this.StopRun()
}

//GetPageCount 翻页计算
func (this *BaseHandler) GetPageCount(count, limit int64) (pagecount int64) {
	if count > 0 && limit > 0 {
		if count%limit == 0 {
			pagecount = count / limit
		} else {
			pagecount = (count / limit) + 1
		}
	}
	return pagecount
}

func (this *BaseHandler) DeleteCookie(key string) {
	if len(key) > 0 {
		this.Ctx.SetCookie(key, "", -10000)
	}
}