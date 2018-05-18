package routers

import (
	"github.com/astaxie/beego"
	"github.com/w2hhda/candy/controllers"
)

func init() {
	beego.Include(&controllers.UserController{},
		&controllers.CandyController{},
		&controllers.RankController{},
		&controllers.RecordController{},
		&controllers.GameController{},
		&controllers.AdminController{})
}
