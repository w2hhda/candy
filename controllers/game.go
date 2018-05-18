package controllers

import (
	"encoding/json"
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type GameController struct {
	BaseController
}

func (c *GameController) URLMapping() {
	c.Mapping("RecordGameData", c.RecordGameData)
}

func (c *GameController) TableName() string {
	return "game"
}

// @router /api/game/record [*]
func (c *GameController) RecordGameData() {

	var request models.GameData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.RetError(errParse)
		return
	}

	beego.Info("request", request)
	valid := validation.Validation{}
	valid.Required(request.Count, "count")
	valid.Required(request.Type, "type")
	valid.Required(request.Addr, "addr")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
		}
		c.RetError(errParams)
		return
	}

	err = models.RecordGameData(request)
	if err != nil {
		c.RetError(errDB)
	} else {
		c.RetSuccess("")
	}

}
