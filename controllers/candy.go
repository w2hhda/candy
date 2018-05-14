package controllers

import "github.com/astaxie/beego"

type CandyController struct {
	BaseController
}

func (c *CandyController) URLMapping() {
	c.Mapping("ListAllCandy", c.ListAllCandy)
}

func (c *CandyController) Prepare()  {
	beego.Error("xxxxxxxxxxxxxxxxxxxxxxx")
}

// @router /api/candy/list [*]
func (c *CandyController) ListAllCandy()  {
	beego.Error("===================")
}

func (c *CandyController) DispatchCandy() {

}