package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/w2hhda/candy/utils"
	"github.com/astaxie/beego"
)

type Candy struct {
	Id             int     `json:"id"`
	AllCount       string  `json:"all_count" orm:"column(all_count)"`
	RemainingCount string  `json:"remaining_count" orm:"column(remaining_count)"`
	TokenAddr      string  `json:"token_addr" orm:"column(token_addr)"`
	CandyLabel     string  `json:"candy_label" orm:"column(candy_label)"`
	CandyType      int     `json:"candy_type" orm:"pk;column(candy_type)"`
	Rate           float64 `json:"rate" orm:"column(rate)"`
}

func (this *Candy) TableName() string {
	return "candy"
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
	var list []Candy
	o := orm.NewOrm()
	page := Page{PageSize: 10}
	count, err := o.QueryTable(this.TableName()).Count()
	if err != nil {
		return page, err
	} else {
		totalPage := count / page.PageSize
		if count > totalPage*page.PageSize {
			totalPage += 1
		}
		page.TotalPage = totalPage
		beego.Info("totalPage=", totalPage)
		if lastPageNumber+1 > totalPage {
			return page, nil
		} else {
			beego.Info("offset=", (lastPageNumber)*page.PageSize)
			page.PageNumber = lastPageNumber + 1
			_, err := o.QueryTable(this.TableName()).Limit(page.PageSize, lastPageNumber*page.PageSize).OrderBy("id").All(&list)
			page.List = list
			return page, err
		}
	}
	return page, err
}

func (this *Candy) DistributionCandy() string {
	o := orm.NewOrm()
	o.QueryTable(this.TableName()).Filter("")
	return utils.RandInt(10, 20)

}

func (this *Candy) DistributionDiamond() string {

	o := orm.NewOrm()
	o.QueryTable(this.TableName()).Filter("")

	return utils.RandInt(1, 10)
}
