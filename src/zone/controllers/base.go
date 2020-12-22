package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gozone/library/enum"
	"gozone/library/logger"
	"gozone/library/util"
	"gozone/src/zone/auth"
	"gozone/src/zone/models"
	"time"
)

const SessionUserKey = "session_user_key"

// gozone基础控制器
type BaseHandler struct {
	Host           string // 访问者的域名
	Body           string // 请求body
	Method         string
	ClientIP       string // 访问者IP
	UserAgent      string
	RequestURL     string
	ActionName     string
	ControllerName string

	IsLogin bool        // 是否登录
	Pager   *Pager      // 分页Pager
	User    models.User // 用户信息
	beego.Controller
}

// 返回数据模板结构
type DataResponse struct {
	Code      enum.ResponseCode `json:"code"`       // 返回码
	Msg       string            `json:"msg"`        // 消息
	StartTime string            `json:"start_time"` // 时间戳
	Body      *Body             `json:"body"`
}

// 具体数据模型
type Body struct {
	Pager *Pager      `json:"pager"`
	Data  interface{} `json:"data"`
}

// 分页Pager
type Pager struct {
	Page      int64 `json:"page"`       //页数
	Limit     int64 `json:"limit"`      //大小
	Offset    int64 `json:"offset"`     //偏移量
	Count     int64 `json:"count"`      //总数
	PageCount int64 `json:"page_count"` //当前页数量
}

func (this *BaseHandler) Prepare() {
	this.Parse()
	this.IsLogin = false
	session := this.GetSession(SessionUserKey)
	if session != nil {
		if user, ok := session.(models.User); ok {
			this.User = user
			userMap, _ := util.Struct2JsonMap(user)
			this.Data["User"] = userMap
			this.IsLogin = true
		}
	}
	cookie := this.Ctx.GetCookie("toggleTheme")

	this.Data["HowToLive"] = util.HowManyToLive(time.Now(), time.Unix(util.BirthDay20201127, 0))
	this.Data["ToggleTheme"] = cookie
	this.Data["IsLogin"] = this.IsLogin
}

// 解析请求信息
func (this *BaseHandler) Parse() {
	this.GetServerParams()
	this.Data["version"] = 0
	this.Data["year"] = time.Now().Year()
	logger.ZoneLogger.Info(fmt.Sprintf("%s | %s | %s | %s", this.ClientIP, this.Ctx.Request.Method, this.Ctx.Request.Host, this.Ctx.Request.RequestURI))
}

// 回复接口
// @param code 回复码
// @param msg 回复消息
// @param args 回复数据
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
		logger.ZoneLogger.Errorf("%v | ERROR: %v", this.RequestURL, err)
	}

	if code == 0 && len(msg) == 0 {
		msg = "success"
	}

	resp := &DataResponse{
		Code:      code,
		Msg:       msg,
		StartTime: time.Now().Format("2006-01-02 15:04:05"),
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

// 翻页计算
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

// 处理控制所需要参数
func (this *BaseHandler) GetServerParams() {
	this.ClientIP = this.GetIP()
	this.Host = this.Ctx.Request.Host
	this.Method = this.Ctx.Request.Method
	this.RequestURL = this.Ctx.Request.RequestURI
	this.Body = string(this.Ctx.Input.RequestBody)
	this.UserAgent = this.Ctx.Input.UserAgent()
	this.ControllerName, this.ActionName = this.GetControllerAndAction()

	//处理翻页
	pager := new(Pager)
	pager.Page, _ = this.GetInt64("page", 1)
	if pager.Page < 1 {
		pager.Page = 1
	}
	pager.Limit, _ = this.GetInt64("limit", 4)
	if pager.Limit < 1 {
		pager.Limit = 4
	}
	pager.Offset = (pager.Page - 1) * pager.Limit
	this.Pager = pager

}

// 获取用户IP
func (this *BaseHandler) GetIP() (ip string) {
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
		ip = "127.0.0.1"
	}
	return
}

// 设置cookie
func (this *BaseHandler) SetCK(key, val string, expireHour int64) {
	if len(key) == 0 || len(val) == 0 {
		return
	}
	val = new(util.XXTEA).EncryptString(val, auth.XXTEKEY)
	this.Ctx.SetCookie(key, val, expireHour*60*60, "/", "", false, true)
}

// 删除cookie
func (this *BaseHandler) DeleteCookie(key string) {
	if len(key) > 0 {
		this.Ctx.SetCookie(key, "", -10000)
	}
}

// 未登录用户重置到登录界面
func (this *BaseHandler) MustLogin() {
	if !this.IsLogin {
		this.Redirect("/", 302)
		return
	}
}
