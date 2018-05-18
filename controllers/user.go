package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/w2hhda/candy/models"

	"strconv"
	"github.com/astaxie/beego/validation"
	"errors"
	"encoding/json"
)

const ADDR = "addr"
const TYPE = "type"

// 一个addr代表一个用户
// 钱包app上的用户可能有多个addr
type UserController struct {
	BaseController
}

func (c *UserController) URLMapping() {
	c.Mapping("GetToken", c.GetToken)
	c.Mapping("SetToken", c.SetToken)
	c.Mapping("ListUserCandy", c.ListUserCandy)
}

// 列出一个钱包用户身上的所有糖果，一个钱包拥有多个地址
// @router /api/token/list [*]
func (c *UserController) ListUserCandy() {

	var request Data
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		beego.Warn(err)
		c.RetError(errParse)
		return
	}

	beego.Info("request=", request)
	list, err := models.ListUserCandyByAddr(request.Addrs)

	var retValue []models.UserToken
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
		value := models.UserToken{
			Addr: addrs, Label: label, Count: tokenCount[label], Rate: tokenRate[label],
		}
		retValue = append(retValue, value)
	}

	if err != nil {
		beego.Warn(err)
		c.RetError(errDB)
	} else {
		c.RetSuccess(retValue)
	}
}

// @router /api/token/get [*]
func (c *UserController) GetToken() {

	token := models.Token{}
	if err := parseAndValidParams(&c.BaseController, &token); err != nil {
		return
	}

	beego.Info("token=", token)

	o := orm.NewOrm()
	err := o.Read(&token, ADDR, TYPE)
	if err != nil {
		beego.Warn(err)
		c.RetError(errParams)
	} else {
		c.RetSuccess(token)
	}
}

// @router /api/token/set [*]
func (c *UserController) SetToken() {

	token := models.Token{}
	if err := parseAndValidParams(&c.BaseController, &token); err != nil {
		return
	}

	beego.Info("token=", token)

	var addCount float64
	var amountErr error
	addCount, amountErr = strconv.ParseFloat(token.Count, 64)
	if amountErr != nil {
		c.RetError(errParams)
		return
	}

	o := orm.NewOrm()
	err := o.Read(&token, ADDR, TYPE)

	if err == nil || err == orm.ErrNoRows {
		dbCount := parseFloat(token.Count) + addCount
		token.Count = parseString(dbCount)
		if err == nil {
			_, err := o.Update(token)
			if err != nil {
				beego.Error(err)
			}
		} else {
			_, err := o.Insert(dbCount)
			if err != nil {
				beego.Error(err)
			}
		}
		c.RetSuccess(new(interface{}))
	} else {
		c.RetError(errDB)
	}

}

func parseAndValidParams(c *BaseController, token *models.Token) error {

	if err := c.ParseForm(token); err != nil {
		c.RetError(errParams)
		return errors.New("params error")
	}

	valid := validation.Validation{}
	if b, err := valid.Valid(token); !b || err != nil {
		if !b {
			for _, err := range valid.Errors {
				beego.Error(err.Key, err.Message)
			}
		}
		c.RetError(errParams)
		return errors.New("params error")
	}

	return nil
}
