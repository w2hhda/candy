package controllers

import (
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"math/big"
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

// @router /api/candy/index [*]
func (c *CandyController) ListAllCandyCountAndGame() {
	act := models.Candy{}
	list, err := act.ListAllCandy()
	if err != nil {
		beego.Warn(err)
		c.RetError(errDB)
	} else {
		game := models.Game{}
		gList, err := game.ListAllGame()
		if err != nil {
			beego.Warn(err)
			c.RetError(errDB)
		}

		var allCount big.Int
		for _, candy := range list {
			if err != nil {
				c.RetError(errDB)
				break
			} else {
				count, _ := big.NewInt(1).SetString(candy.AllCount, 10)
				allCount = *big.NewInt(1).Add(&allCount, count)
			}
		}

		c.RetSuccess(map[string]interface{}{
			"all_candy_count": allCount.String(),
			"all_game_list":   gList,
		})
	}
}
