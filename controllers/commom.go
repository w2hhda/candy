package controllers

var (
	successReturn = &Response{0, "success", new(interface{})}
	errParams     = &Response{10001, "输入的参数不正确", new(interface{})}
	errDB         = &Response{10002, "数据库错误", new(interface{})}
)

func (base *BaseController) RetError(e *Response) {
	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Data["json"] = e
	base.ServeJSON()
	base.StopRun()
}

func (base *BaseController) RetSuccess(data interface{}) {
	successReturn.Value = data
	base.Data["json"] = successReturn
	base.ServeJSON()
}
