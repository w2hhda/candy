package models

import (
	"github.com/astaxie/beego"
)

type RankInfo struct {
	Addr  string `json:"addr"`
	Count string `json:"count"`
	Value string `json:"value"`
}

func Rank(lastPageNumber int64) (Page, error) {

	page, err := countPage(UserCandyTableName(), lastPageNumber)
	var rankInfo []RankInfo
	db := DB()
	stmt, _ := db.Prepare("SELECT SUM(token.count) AS count, addr FROM token GROUP BY addr ORDER BY count DESC LIMIT ? OFFSET ?;")
	rows, fErr := stmt.Query(page.PageSize, lastPageNumber*page.PageSize)
	if fErr != nil {
		beego.Warn(fErr)
		return page, nil
	}

	for rows.Next() {
		value := RankInfo{}
		rows.Scan(&value.Count, &value.Addr)
		beego.Info(value.Addr)
		//计算价格
		token, _ := ListUserCandyByAddr([]string{value.Addr})
		var price float64
		for _, t := range token {
			price += t.Candy.Rate * parseFloat(t.Count)
		}
		beego.Info("==>>", price)
		value.Value = parseString(price)
		rankInfo = append(rankInfo, value)
	}
	page.List = rankInfo
	page.PageSize = int64(len(rankInfo))
	return page, err
}
