package controllers

import (
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
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
	client := models.Redis()
	var page models.Page
	err = client.Get(models.GetRankCacheKey(request.PageNumber)).Scan(&page)
	if err == nil {
		beego.Warn("hit cache")
		c.RetSuccess(page)
	} else {
		beego.Warn("miss cache")
		c.getDBRank(client, &request)
	}
}

func (c *RankController) getDBRank(client *redis.Client, request *RequestData) {
	values, err := models.Rank(request.PageNumber)
	if err != nil {
		beego.Warn(err)
		c.RetError(errDB)
	}
	by, err := json.Marshal(values)
	if err == nil {
		status, err := client.Set(models.GetRankCacheKey(request.PageNumber), by, time.Hour*12).Result()
		beego.Warn(status, err)
	}
	c.RetSuccess(values)
}
