package mail

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
)

var password = beego.AppConfig.String("mail:password")

// 发送一个新邮件
// @param email 接收方邮箱
// @param code 注册验证码
// @return 错误信息
func SendMail(email string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "522240909@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "[pannnkee.com] Please Confirm Your Verify Code")
	m.SetBody("text/html", fmt.Sprintf("验证码:%s", code))
	d := gomail.NewDialer("smtp.qq.com", 465, "522240909@qq.com", "arrzshmjevcscaig")
	err := d.DialAndSend(m)
	return err
}
