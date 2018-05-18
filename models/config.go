package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	"strconv"
)

type Config struct {
	Id    int         `orm:"column(id)"`
	Key   string      `orm:"column(key)"`
	Value interface{} `orm:"column(value)"`
}

func (c *Config) TableName() string {
	return "app_config"
}

func (c *Config) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(c.TableName())
}

func (c *Config) String() string {
	var config Config
	err := c.Query().Filter("key", c.Key).One(&config);
	if err != nil {
		beego.Warn(err)
		return ""
	} else {
		return config.Value.(string)
	}
}

func (c *Config) Int() int64 {
	var config Config
	err := c.Query().Filter("key", c.Key).One(&config);
	if err != nil {
		beego.Warn(err)
		return -1
	} else {
		value, err := strconv.ParseInt(config.Value.(string), 10, 64)
		if err != nil {
			beego.Warn(err)
			return -1
		}
		return value
	}
}

func (c *Config) Set() {
	o := orm.NewOrm()
	b := o.QueryTable(c.TableName()).Filter("key", c.Key).Exist()
	if b {
		re, err := o.Update(c, "value")
		beego.Info("Update ", c, " Return ", re)
		if err != nil {
			beego.Warn(err)
		}
	} else {
		re, err := o.Insert(c)
		beego.Info("Insert ", c, " Return ", re)
		if err != nil {
			beego.Warn(err)
		}
	}

}
