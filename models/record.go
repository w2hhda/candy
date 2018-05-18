package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Record struct {
	Id       int    `json:"id" orm:"column(id)"`
	Addr     string `json:"addr" orm:"column(addr)"`
	Count    string `json:"count" orm:"column(count)"`
	Candy    *Candy `json:"candy" orm:"rel(one)"`
	CreateAt string `json:"create_at" orm:"column(create_at)"`
	Game     *Game  `json:"game" orm:"rel(one)"`
}

func (r *Record) TableName() string {
	return "record"
}

func (r *Record) ListRecordByAddr(lastPageNumber int64, addrs []string) (Page, error) {

	var list []Record
	o := orm.NewOrm()
	page := Page{PageSize: 10}
	count, err := o.QueryTable(r.TableName()).Count()
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
			_, err := o.QueryTable(r.TableName()).Filter("addr__in", addrs).
				RelatedSel().Limit(page.PageSize, lastPageNumber*page.PageSize).OrderBy("-create_at").All(&list)
			page.List = list
			return page, err
		}
	}
	return page, err

}
