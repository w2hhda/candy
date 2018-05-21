package controllers

import (
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"time"
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

	client := models.Redis()
	var indexInfo models.IndexInfo
	err := client.Get(models.GetIndexCacheKey()).Scan(&indexInfo)
	if err != nil {
		indexInfo, err := models.ListIndex()
		if err != nil {
			c.RetError(errDB)
		} else {
			beego.Info(indexInfo)
			by, err := json.Marshal(indexInfo)
			if err == nil {
				status, err := client.Set(models.GetIndexCacheKey(), by, 12*time.Hour).Result()
				beego.Warn("status", status, err)
			}
			c.RetSuccess(indexInfo)
		}
	} else {
		c.RetSuccess(indexInfo)
	}

}
