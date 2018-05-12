package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	user := beego.AppConfig.String("sqluser")
	password := beego.AppConfig.String("sqlpass")
	dbName := beego.AppConfig.String("sqldb")
	orm.RegisterDataBase("default", "mysql", user+":"+password+"@/"+dbName+"?charset=utf8")
}

