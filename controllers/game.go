package controllers

import (
	"encoding/json"
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"math/big"
	"sync"
)

type GameController struct {
	BaseController
}

func (c *GameController) URLMapping() {
	c.Mapping("RecordGameData", c.RecordGameData)
	c.Mapping("GameStart", c.GameStart)
	c.Mapping("GameOver", c.GameOver)
}

func (c *GameController) TableName() string {
	return "game"
}

func GameCandyTableName() string {
	return "gamecandy"
}

// @router /api/game/start [*]
func (c *GameController) GameStart() {
	var request models.GameRunningData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		beego.Info("GameStart", err)
		c.RetError(errParse)
		return
	}

	beego.Info("request", request)
	valid := validation.Validation{}
	valid.Required(request.GameId, "game_id")
	valid.Required(request.GameFieldId, "game_field_id")
	result := valid.Required(request.GamePlayer, "game_player")

	if result.Ok {
		for _, player := range request.GamePlayer {
			beego.Info(player)
			valid.Required(player.Addr, "p_addr")
		}
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
		}
		c.RetError(errParams)
		return
	}

	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	// 1. 分配糖果，查询糖果是否 > 0
	normal, diamond, err := models.ListUsefulCandy(len(request.GamePlayer), 3)
	if err != nil {
		res := &models.Response{Code: 10006, Message: err.Error(), Value: ""}
		c.RetError(res)
		return
	}

	// 2. 游戏开始创建
	ns, ds, err := models.ReadOrInsert(normal, diamond, request)
	if err != nil {
		c.RetError(errDB)
		return
	}

	ngsd := models.GameStartData{
		CandyType:  normal.CandyType,
		CandyLabel: normal.Alias,
		CandyCount: ns,
	}
	if diamond.Id > 0 {
		c.RetSuccess([2]models.GameStartData{
			ngsd,
			{
				CandyType:  diamond.CandyType,
				CandyLabel: diamond.Alias,
				CandyCount: ds,
			},
		})
	} else {
		c.RetSuccess([1]models.GameStartData{
			ngsd,
		})
	}

}

// @router /api/game/over [*]
func (c *GameController) GameOver() {
	var request models.GameRunningData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		beego.Info("GameStart", err)
		c.RetError(errParse)
		return
	}

	beego.Info("request", request)
	valid := validation.Validation{}
	valid.Required(request.GameId, "game_id")
	valid.Required(request.GameFieldId, "game_field_id")
	result := valid.Required(request.GamePlayer, "game_player")

	if result.Ok {
		for _, player := range request.GamePlayer {
			beego.Info(player)
			valid.Required(player.Addr, "p_addr")
			valid.Required(player.Score, "p_score")
		}
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
		}
		c.RetError(errParams)
		return
	}
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
