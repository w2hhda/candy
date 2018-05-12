package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["candy/controllers:TokenController"] = append(beego.GlobalControllerRouter["candy/controllers:TokenController"],
		beego.ControllerComments{
			Method: "GetToken",
			Router: `/api/token/get`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["candy/controllers:TokenController"] = append(beego.GlobalControllerRouter["candy/controllers:TokenController"],
		beego.ControllerComments{
			Method: "ListToken",
			Router: `/api/token/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["candy/controllers:TokenController"] = append(beego.GlobalControllerRouter["candy/controllers:TokenController"],
		beego.ControllerComments{
			Method: "SetToken",
			Router: `/api/token/set`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
