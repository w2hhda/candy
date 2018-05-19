package models

import (
	"strconv"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"math/big"
)

type Request struct {
	AppVersion string `json:"app_version"`
	OsVersion  string `json:"os_version"`
	Imei       string `json:"imei"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

type Page struct {
	PageNumber int64       `json:"page_number"`
	PageSize   int64       `json:"page_size"`
	TotalPage  int64       `json:"total_page"`
	TotalCount int64       `json:"total_count"`
	List       interface{} `json:"list"`
}

type LayuiPageResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

type UserToken struct {
	Addr  []string `json:"addr"`
	Label string   `json:"label"`
	Count big.Int  `json:"count"`
	Icon  string   `json:"icon"`
	Rate  float64  `json:"rate"`
}

func parseFloat(input string) float64 {
	out, _ := strconv.ParseFloat(input, 64)
	return out
}

func parseString(input float64) string {
	if input == 0 {
		return "0"
	}
	return strconv.FormatFloat(input, 'E', -1, 64)
}

func countPage(tableName string, lastPageNumber int64) (Page, error) {
	o := orm.NewOrm()
	page := Page{PageSize: 10}
	count, err := o.QueryTable(tableName).Count()
	page.TotalCount = count;
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
			beego.Info("offset=", lastPageNumber*page.PageSize)
			page.PageNumber = lastPageNumber + 1
			return page, err
		}
	}
}
