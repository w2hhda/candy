package controllers

import (
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type RankController struct {
	BaseController
}

func (c *RankController) URLMapping() {
	c.Mapping("Rank", c.Rank)
}

// @router /api/rank [*]
func (c *RankController) Rank() {

	var request RequestData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		beego.Warn(err)
		c.RetError(errParse)
		return
	}

	beego.Info("request", request)

	values, err := models.Rank(request.PageNumber)
	if err != nil {
		beego.Warn(err)
		c.RetError(errDB)
	}
	c.RetSuccess(values)

}
