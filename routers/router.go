package routers

import (
	"github.com/astaxie/beego"
	"github.com/w2hhda/candy/controllers"
)

func init() {
	beego.Include(&controllers.TokenController{}, &controllers.CandyController{})
}
