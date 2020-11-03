package conn

import (
	"gozone/library/config"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
)

var DataBases map[string]*gorm.DB
var DefaultDBName  = config.GetConfigStr("mysql_dbname", "zone")

func init() {
	DataBases = make(map[string]*gorm.DB)
	initDataBase(DefaultDBName)
}

func initDataBase(dbName string)  {
	var orm *gorm.DB
	var err error

	user := config.GetConfigStr(dbName + "::user", "root")
	password := config.GetConfigStr(dbName + "::password", "root123")
	host := config.GetConfigStr(dbName + "::host", "127.0.0.1")
	port := beego.AppConfig.DefaultInt(dbName+"::port", 3306)
	if len(user) == 0 || len(password) == 0 || len(host) == 0 || port == 0 {
		panic("数据库配置信息失败")
	}

	// charset = utf8&
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName) + "?charset=utf8mb4&parseTime=true"
	for orm, err = gorm.Open("mysql", connStr); err != nil; {
		logs.Error("[mysql连接异常，正在重试:]", err.Error(), connStr)
		time.Sleep(5*time.Second)
		orm, err = gorm.Open("mysql", connStr)
	}

	if config.GetConfigStr("runmode", "") != "prod" {
		orm.LogMode(true)
	}

	orm.CommonDB()
	orm.DB().SetMaxOpenConns(1000)
	orm.DB().SetMaxIdleConns(10)
	orm.DB().SetConnMaxLifetime(time.Hour)
	DataBases[dbName] = orm

	logs.Info(fmt.Sprintf("连接mysql成功：%v", connStr))
}

func GetORMByName (dbName string) *gorm.DB {
	db := DataBases[dbName]
	if strings.ToLower(config.GetConfigStr("runmode", "dev")) == "prod" {
		db.LogMode(false)
	}
	return db
}
