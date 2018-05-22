package controllers

import (
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type CandyController struct {
	BaseController
}

func (c *CandyController) URLMapping() {
	c.Mapping("ListAllCandyCountAndGame", c.ListAllCandyCountAndGame)
	c.Mapping("ListCandyPage", c.ListCandyPage)
}

// @router /api/candy/list [*]
func (c *CandyController) ListCandyPage() {

	var request RequestData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.RetError(errParse)
		return
	}

	beego.Info("request=", request)

	act := models.Candy{}
	page, err := act.ListCandyPage(request.PageNumber)
	if err != nil {
		beego.Warn(err)
		c.RetError(errDB)
	} else {
		c.RetSuccess(page)
	}
}

// 钱包app 应用tab 显示的内容
// @router /api/candy/index [*]
func (c *CandyController) ListAllCandyCountAndGame() {
	indexInfo, err := models.ListIndex()
	if err != nil {
		c.RetError(errDB)
	} else {
		beego.Info(indexInfo)
		c.RetSuccess(indexInfo)
	}
}
