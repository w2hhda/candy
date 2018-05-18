package models

import "github.com/astaxie/beego/orm"

type Game struct {
	Id     int    `json:"id" orm:"column(id)"`
	Link   string `json:"url" orm:"column(link)"`
	Sort   int    `json:"sort" orm:"column(sort)"`
	Status int    `json:"status" orm:"column(status)"`
	Icon   string `json:"icon" orm:"column(icon)"`
	Name   string `json:"name" orm:"column(name)"`
}

func (this *Game) TableName() string {
	return "game"
}

func (this *Game) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName())
}

func (this *Game) ListAllGame() ([]Game, error) {
	var list []Game
	_, err := this.Query().All(&list)
	return list, err
}
