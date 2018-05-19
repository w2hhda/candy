package models

import (
	"github.com/astaxie/beego/orm"
)

type Candy struct {
	Id             int     `json:"id"`
	AllCount       string  `json:"all_count" orm:"column(all_count)"`
	RemainingCount string  `json:"remaining_count" orm:"column(remaining_count)"`
	TokenAddr      string  `json:"token_addr" orm:"column(token_addr)"`
	CandyLabel     string  `json:"candy_label" orm:"column(candy_label)"`
	CandyType      int     `json:"candy_type" orm:"pk;column(candy_type)"`
	Rate           float64 `json:"rate" orm:"column(rate)"`
	Decimal        int     `json:"decimal" orm:"column(decimal)"`
}

func CandyTableName() string {
	return "candy"
}

func (this *Candy) TableName() string {
	return CandyTableName()
}

func (this *Candy) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName())
}

func (this *Candy) ListAllCandy() ([]Candy, error) {
	var list []Candy
	_, err := this.Query().All(&list)
	return list, err
}

func (this *Candy) ListCandyPage(lastPageNumber int64) (Page, error) {
	page, err := countPage(this.TableName(), lastPageNumber)
	if err == nil {
		var list []Candy
		_, err = orm.NewOrm().QueryTable(this.TableName()).Limit(page.PageSize, lastPageNumber*page.PageSize).OrderBy("id").All(&list)
		page.List = list
		page.PageSize = int64(len(list))
	}
	return page, err
}
