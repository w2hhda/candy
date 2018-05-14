package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:CandyController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:CandyController"],
		beego.ControllerComments{
			Method: "ListAllCandy",
			Router: `/api/candy/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:TokenController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:TokenController"],
		beego.ControllerComments{
			Method: "GetToken",
			Router: `/api/token/get`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:TokenController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:TokenController"],
		beego.ControllerComments{
			Method: "ListToken",
			Router: `/api/token/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:TokenController"] = append(beego.GlobalControllerRouter["github.com/w2hhda/candy/controllers:TokenController"],
		beego.ControllerComments{
			Method: "SetToken",
			Router: `/api/token/set`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
