package models

import "github.com/astaxie/beego/orm"

type Token struct {
	Id    int    `orm:"column(id)" json:"id" form:"-"`
	Addr  string `orm:"column(addr)" json:"addr" form:"addr" valid:"Required"`
	Name  string `orm:"column(name)" json:"name" form:"name" valid:"MaxSize(40)"`
	Type  string `orm:"column(type)" json:"type" form:"type" valid:"Required;MaxSize(10)"`
	Count string `orm:"column(count)" json:"count" form:"count" valid:"Numeric"`
}

func init() {
	orm.RegisterModel(new(Token))
}
