package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Candy",
			Router: `/admin/candy`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"],
		beego.ControllerComments{
			Method: "ListCandy",
			Router: `/admin/candy/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"],
		beego.ControllerComments{
			Method: "DisableUser",
			Router: `/admin/disable`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/admin/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"],
		beego.ControllerComments{
			Method: "User",
			Router: `/admin/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:AdminController"],
		beego.ControllerComments{
			Method: "ListUser",
			Router: `/admin/user/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:CandyController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:CandyController"],
		beego.ControllerComments{
			Method: "ListAllCandyCountAndGame",
			Router: `/api/candy/index`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:CandyController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:CandyController"],
		beego.ControllerComments{
			Method: "ListCandyPage",
			Router: `/api/candy/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:GameController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:GameController"],
		beego.ControllerComments{
			Method: "GameOver",
			Router: `/api/game/over`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:GameController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:GameController"],
		beego.ControllerComments{
			Method: "GameStart",
			Router: `/api/game/start`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:RankController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:RankController"],
		beego.ControllerComments{
			Method: "Rank",
			Router: `/api/rank`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:RecordController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:RecordController"],
		beego.ControllerComments{
			Method: "Record",
			Router: `/api/record/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:UserController"],
		beego.ControllerComments{
			Method: "ListUserCandy",
			Router: `/api/token/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
