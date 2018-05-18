package controllers

import (
	"github.com/w2hhda/candy/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type RecordController struct {
	BaseController
}

type RecordData struct {
	models.Request
	PageNumber int64    `json:"page_number"`
	Addrs      []string `json:"addrs"`
}

func (c *RecordController) URLMapping() {
	c.Mapping("Record", c.Record)
}

// @router /api/record/list [*]
func (c *RecordController) Record() {

	var request RecordData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.RetError(errParse)
		return
	}

	beego.Info("request=", request)

	record := models.Record{}
	page, err := record.ListRecordByAddr(request.PageNumber, request.Addrs)
	if err != nil {
		beego.Warn(err)
		c.RetError(errDB)
		return
	}

	c.RetSuccess(page)
	return

}
