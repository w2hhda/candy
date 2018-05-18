package controllers

import (
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"strconv"
	"encoding/json"
)

type CandyController struct {
	BaseController
}

func (c *CandyController) URLMapping() {
	c.Mapping("ListAllCandyCountAndGame", c.ListAllCandyCountAndGame)
	c.Mapping("ListCandyPage", c.ListCandyPage)
	c.Mapping("DistributionCandy", c.DistributionCandy)
}

// @router /api/candy/list [*]
func (c *CandyController) ListCandyPage() {

	var request RecordData
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

		var allCount float64
		for _, candy := range list {
			count, err := strconv.ParseFloat(candy.AllCount, 64)
			if err != nil {
				c.RetError(errDB)
				break
			} else {
				allCount += count
			}
		}

		c.RetSuccess(map[string]interface{}{
			"all_candy_count": parseString(allCount),
			"all_game_list":   gList,
		})
	}
}

// @router /api/candy/distribution [*]
func (c *CandyController) DistributionCandy() {

}
