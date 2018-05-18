package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type User struct {
	Id       int    `json:"id" orm:"column(id)"`
	Addr     string `json:"addr" orm:"column(addr)"`
	CreateAt string `json:"create_at" orm:"column(create_at)"`
	Status   int    `json:"status" orm:"column(status)"`
	Name     string `json:"name" orm:"column(name)"`
}

type Token struct {
	Id       int    `orm:"column(id)" json:"id" form:"-"`
	Addr     string `orm:"pk;column(addr)" json:"addr" form:"addr" valid:"Required"`
	Count    string `orm:"column(count)" json:"count" form:"count" valid:"Numeric"`
	Candy    *Candy `orm:"rel(one)"`
	UpdateAt string `json:"update_at" orm:"column(update_at)"`
}

func (t *Token) TableName() string {
	return "token"
}

func ListCandyPage(lastPageNumber int64, limit int64) (Page, error) {

	//1. 从user表中获取 addr
	page, err := ListUserPage(lastPageNumber, limit)

	//2. 通过user获取糖果的总量
	var (
		addrs    []string
		retValue []UserToken
	)
	for _, user := range page.List.([]User) {
		addrs = append(addrs, user.Addr)
	}

	beego.Info("user", addrs)
	list, _ := ListTokenByAddr(addrs)

	tokenCount := make(map[string]string)
	tokenAddr := make(map[string][]string)
	tokenRate := make(map[string]float64)
	for _, token := range list {
		//糖果类型
		candyLabel := token.Candy.CandyLabel
		//糖果数量
		count := parseFloat(tokenCount[candyLabel]) + parseFloat(token.Count)
		tokenCount[candyLabel] = parseString(count)
		//糖果地址
		tokenAddr[candyLabel] = append(tokenAddr[candyLabel], token.Addr)
		//糖果价格
		tokenRate[candyLabel] = token.Candy.Rate
	}

	for label, addrs := range tokenAddr {
		value := UserToken{
			Addr: addrs, Label: label, Count: tokenCount[label], Rate: tokenRate[label],
		}
		retValue = append(retValue, value)
	}

	page.List = retValue
	page.PageSize = int64(len(list))
	beego.Warn(list)

	return page, err

}

//后台获取所有用户信息，分页查询, 调整位置
func ListUserPage(lastPageNumber int64, limit int64) (Page, error) {
	var list []User
	userTableName := "user"
	page, err := countPage(userTableName, lastPageNumber)
	if err != nil {
		return page, err
	}
	_, err = orm.NewOrm().QueryTable(userTableName).
		Limit(limit, lastPageNumber*page.PageSize).OrderBy("id").All(&list)
	page.List = list
	page.PageSize = int64(len(list))
	beego.Info(list)
	return page, err
}

func ListTokenByAddr(addrs []string) ([]Token, error) {
	var token []Token
	_, err := orm.NewOrm().QueryTable("token").Filter("addr__in", addrs).RelatedSel().All(&token)
	beego.Info(token)
	return token, err

}
