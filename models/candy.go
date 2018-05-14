package models

import "github.com/astaxie/beego/orm"

type AllCandyToken struct {
	Id             int    `orm:"column(id)"`
	Label          string `json:"label" orm:"column(lable)"`
	Type           string `json:"type" orm:"column(type)"`
	AllCount       int    `json:"all_count" orm:"column(all_count)"`
	RemainingCount int    `json:"remaining_count" orm:"column(remaining_count)"`
}

type Candy struct {
	Label string `json:"label"`
	Type  int    `json:"type"`
	Count string `json:"count"`
}

func init() {
	orm.RegisterModel(AllCandyToken{})
}

func (this *AllCandyToken) TableName() string {
	return "all_token"
}

func (this *AllCandyToken) DistributionCandy() {

	o := orm.NewOrm();
	o.QueryTable(this.TableName()).Filter("")

}
