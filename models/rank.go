package models

import (
	"github.com/astaxie/beego"
	"math/big"
)

type RankInfo struct {
	Addr  string `json:"addr"`
	Count string `json:"count"`
	Value string `json:"value"`
}

func Rank(lastPageNumber int64) (Page, error) {

	page, err := countPage(TokenTableName(), lastPageNumber)
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
		var price big.Float
		for _, t := range token {
			f, _, _ := new(big.Float).Parse(t.Count, 10)
			r, _, _ := new(big.Float).Parse(parseString(t.Candy.Rate), 10)
			price = *new(big.Float).Add(new(big.Float).Mul(r, f), &price)
		}
		value.Value = price.String()
		rankInfo = append(rankInfo, value)
	}
	page.List = rankInfo
	page.PageSize = int64(len(rankInfo))

	return page, err
}
