package models

import (
	"github.com/astaxie/beego"
	"database/sql"
	"github.com/astaxie/beego/orm"
)

type RankInfo struct {
	Addr  string `json:"addr"`
	Count string `json:"count"`
	Value string `json:"value"`
}

func Rank(lastPageNumber int64) (Page, error) {

	page := Page{PageSize: 10}

	count, err := orm.NewOrm().QueryTable("token").Count()
	if err != nil {
		return page, err
	}

	totalPage := count / page.PageSize
	if count > totalPage*page.PageSize {
		totalPage += 1
	}
	page.TotalPage = totalPage
	beego.Info("totalPage=", totalPage)
	if lastPageNumber+1 > totalPage {
		return page, nil
	} else {
		beego.Info("offset=", lastPageNumber*page.PageSize)
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
			token, _ := ListTokenByAddr([]string{value.Addr})
			var price float64
			for _, t := range token {
				price += t.Candy.Rate * parseFloat(t.Count)
			}
			beego.Info("==>>", price)
			value.Value = parseString(price)
			rankInfo = append(rankInfo, value)
		}

		page.List = rankInfo

	}

	return page, nil
}

// DB function
func DB() *sql.DB {

	user := beego.AppConfig.String("sqluser")
	password := beego.AppConfig.String("sqlpass")
	host := beego.AppConfig.String("host")
	dbName := beego.AppConfig.String("sqldb")

	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+dbName)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
