package controllers

import (
	"github.com/astaxie/beego"
	"github.com/w2hhda/candy/models"

	"github.com/astaxie/beego/validation"
	"errors"
	"encoding/json"
	"math/big"
)

const ADDR = "addr"
const TYPE = "type"

// 一个addr代表一个用户
// 钱包app上的用户可能有多个addr
type UserController struct {
	BaseController
}

func (c *UserController) URLMapping() {
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
	tokenCount := make(map[string]big.Int)
	tokenAddr := make(map[string][]string)
	tokenRate := make(map[string]float64)
	for _, token := range list {
		//糖果类型
		candyLabel := token.Candy.CandyLabel
		//糖果数量
		candyCount := tokenCount[candyLabel]
		tCount, _ := big.NewInt(1).SetString(token.Count, 10)
		count := big.NewInt(1).Add(&candyCount, tCount)
		tokenCount[candyLabel] = *count
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
