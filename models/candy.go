package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"math/big"
	"errors"
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
	Alias          string  `json:"alias" orm:"column(candy_alias)"`
	Average        int     `json:"average" orm:"column(average)"`
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

func ListUsefulCandy(candyCount int, diamondCount int) (Candy, Candy, error) {
	var (
		list    []Candy
		normal  Candy
		diamond Candy
	)
	_, err := orm.NewOrm().QueryTable(CandyTableName()).All(&list)
	if err == nil {
		//获取普通糖果
		for _, candy := range list {
			beego.Info("candy", candy)
			if candy.CandyType < 10000 {
				remainingCount, _ := new(big.Int).SetString(candy.RemainingCount, 10)
				requireCount := int64(candy.Average * candyCount)
				sub := remainingCount.Cmp(big.NewInt(requireCount))
				if sub != -1 { //足够分配
					normal = candy
					break
				}
			}
		}

		//普通糖果不足就不执行游戏了
		if normal.Id <= 0 {
			return normal, diamond, errors.New("糖果不足")
		}

		//获取钻石
		for _, candy := range list {
			beego.Info("candy", candy)
			if candy.CandyType >= 10000 {
				remainingCount, _ := new(big.Int).SetString(candy.RemainingCount, 10)
				requireCount := int64(candy.Average * diamondCount)
				sub := remainingCount.Cmp(big.NewInt(requireCount))
				if sub != -1 { //足够分配
					diamond = candy
					break
				}
			}
		}
	}
	return normal, diamond, err
}
