package controllers

import "github.com/w2hhda/candy/models"

var (
	successReturn = &models.Response{0, "success", new(interface{})}
	errParams     = &models.Response{10001, "输入的参数不正确", new(interface{})}
	errDB         = &models.Response{10002, "数据库错误", new(interface{})}
	errParse      = &models.Response{10003, "数据解析失败", new(interface{})}
)

func (base *BaseController) RetError(e *models.Response) {
	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Data["json"] = e
	base.ServeJSON()
	base.StopRun()
}

func (base *BaseController) RetSuccess(data interface{}) {
	successReturn.Value = data
	base.Data["json"] = successReturn
	base.ServeJSON()
	base.StopRun()
}

func (base *BaseController) RetLayuiPage(count int64, data interface{}) {
	response := models.LayuiPageResponse{
		Code: 0, Msg: "success", Count: count, Data: data,
	}
	base.Data["json"] = response
	base.ServeJSON()
	base.StopRun()
}
