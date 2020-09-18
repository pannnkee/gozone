package controllers

import (
	"Gozone/library/enum"
	"Gozone/library/util"
	"Gozone/src/zone/auth"
	"Gozone/src/zone/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	IsLogin        bool
	User           models.User
}

const SESSION_USER_KEY = "session_user_key"

// DataResponse 返回数据模型
type DataResponse struct {
	Code  enum.ResponseCode `json:"code"`
	Msg   string            `json:"msg"`
	Stime string            `json:"stime"`
	Body  *Body             `json:"body"`
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

func (this *BaseHandler) Prepare() {
	this.Parse()
	this.IsLogin = false
	session := this.GetSession(SESSION_USER_KEY)
	if session != nil {
		if user, ok := session.(models.User); ok {
			this.User = user
			userMap, _ := util.Struct2JsonMap(user)
			this.Data["User"] = userMap
			this.IsLogin = true
		}
	}

	cookie := this.Ctx.GetCookie("toggleTheme")
	this.Data["ToggleTheme"] = cookie
	this.Data["IsLogin"] = this.IsLogin
}

//Parse Parse
func (this *BaseHandler) Parse() {
	this.GetServerParams()
	this.Data["version"] = 0
	this.Data["ZoneName"] = "pannnkee zone"
	this.Data["year"] = time.Now().Year()
	logs.Info(fmt.Sprintf("%s | %s | %s | %s", this.ClientIP, this.Ctx.Request.Method, this.Ctx.Request.Host, this.Ctx.Request.RequestURI))
}

func (this *BaseHandler) Response(code enum.ResponseCode, msg string, args ...interface{}) {
	var (
		err   error
		pager *Pager
		data  interface{}
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
		Code:  code,
		Msg:   msg,
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

//GetServerParams GetServerParams
func (this *BaseHandler) GetServerParams() {
	this.ClientIP = this.GetIP()
	this.Host = this.Ctx.Request.Host
	//http method
	this.Method = this.Ctx.Request.Method
	// request url
	this.RequestURL = this.Ctx.Request.RequestURI
	//httpbody
	this.Body = string(this.Ctx.Input.RequestBody)
	//controllname

	this.UserAgent = this.Ctx.Input.UserAgent()

	this.ControllerName, this.ActionName = this.GetControllerAndAction()

	//处理翻页
	pager := new(Pager)
	pager.Page, _ = this.GetInt64("page", 1)
	if pager.Page < 1 {
		pager.Page = 1
	}
	pager.Limit, _ = this.GetInt64("limit", 10)
	if pager.Limit < 1 {
		pager.Limit = 10
	}
	pager.Offset = (pager.Page - 1) * pager.Limit
	this.Pager = pager

}

//GetIP 获取用户IP
func (this *BaseHandler) GetIP() (ip string) {
	//适应于API取H5/PC时
	ip = this.Ctx.Request.Header.Get("X-Original-Forwarded-For")
	if len(ip) == 0 {
		ip = this.Ctx.Request.Header.Get("Remote-Host")
	}
	if len(ip) == 0 {
		ip = this.Ctx.Request.Header.Get("X-Real-IP")
	}
	if len(ip) == 0 {
		ip = this.Ctx.Input.IP()
	}
	if len(ip) == 0 {
		ip = "10.0.0.2"
	}
	return
}

//SetCK SetCK
func (this *BaseHandler) SetCK(key, val string, expireHour int64) {
	if len(key) == 0 || len(val) == 0 {
		return
	}
	val = new(util.XXTEA).EncryptString(val, auth.XXTEKEY)
	this.Ctx.SetCookie(key, val, expireHour*60*60, "/", "", false, true)
}

func (this *BaseHandler) DeleteCookie(key string) {
	if len(key) > 0 {
		this.Ctx.SetCookie(key, "", -10000)
	}
}

func (this *BaseHandler) MustLogin() {
	if !this.IsLogin {
		//this.Response(enum.DefaultError, "用户未登录")
		this.Redirect("/", 302)
		return
	}
}
