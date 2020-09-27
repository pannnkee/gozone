package models

import "Gozone/library/conn"

type Log struct {
	ID           int64    `json:"id" gorm:"column:id"`
	Ip           string `json:"ip" gorm:"column:ip"`               // 登录IP
	UserID       int64    `json:"user_id" gorm:"column:user_id"`     // 用户UserID
	UserName     string `json:"user_name" gorm:"column:user_name"` // 用户名
	LoginTime    int64    `json:"login_time" gorm:"column:login_time"`
	LoginTimeStr string `json:"login_time_str" gorm:"column:login_time_str"`
}

func (this *Log) TableName() string {
	return "log"
}

func (this *Log) AddLoginLog() (err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	return db.Create(&this).Error
}
