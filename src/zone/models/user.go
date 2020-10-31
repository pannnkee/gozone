package models

// 用户信息
type User struct {
	Id             int64  `gorm:"column:id" json:"id"`
	UserName       string `gorm:"column:user_name" json:"user_name"`
	Email          string `gorm:"column:email" json:"email"`
	Mobile         string `gorm:"column:mobile" json:"mobile"`
	Avatar         string `gorm:"column:avatar" json:"avatar"`
	PassWord       string `gorm:"column:password" json:"password"`
	RepeatPassword string `gorm:"-" json:"repeat_password"`
	Status         int64  `gorm:"column:status" json:"status"`
	LoginTimes     int64  `gorm:"column:login_time" json:"login_time"`
	CreatedTime    int64  `gorm:"column:created_time" json:"created_time"`
	UpdateTime     int64  `gorm:"column:update_time" json:"update_time"`
}

func (this *User) TableName() string {
	return "user"
}
