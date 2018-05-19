package controllers

import (
	"encoding/json"
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"math/big"
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
		beego.Info("RecordGameData", err)
		c.RetError(errParse)
		return
	}

	beego.Info("request", request)
	valid := validation.Validation{}
	valid.Required(request.Count, "count")
	valid.Required(request.Type, "type")
	valid.Required(request.Addr, "addr")

	_, b := new(big.Int).SetString(request.Count, 10)
	if !b {
		c.RetError(errParams)
		return
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
		}
		c.RetError(errParams)
		return
	}

	err = models.RecordGameData(request)
	if err != nil {
		c.RetError(&models.Response{10004, err.Error(), new(interface{})})
	} else {
		c.RetSuccess("")
	}

}
