package routers

import (
	"candy/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.TokenController{})
}
