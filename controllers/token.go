package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"candy/models"

	"strconv"
	"github.com/astaxie/beego/validation"
	"errors"
)

const ADDR = "addr"
const TYPE = "type"

type TokenController struct {
	BaseController
}

func (c *TokenController) URLMapping() {
	c.Mapping("GetToken", c.GetToken)
	c.Mapping("SetToken", c.SetToken)
	c.Mapping("ListToken", c.ListToken)
}

// @router /api/token/list [*]
func (c *TokenController) ListToken()  {

}

// @router /api/token/get [*]
func (c *TokenController) GetToken() {

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
func (c *TokenController) SetToken() {

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

func parseString(input float64) string {
	return strconv.FormatFloat(input, 'E', -1, 64)
}

func parseFloat(input string) float64 {
	out, _ := strconv.ParseFloat(input, 64)
	return out
}
