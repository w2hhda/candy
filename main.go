package main

import (
	_ "candy/routers"
	"github.com/astaxie/beego"
	_ "candy/models"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetLogger("file", `{"filename":"logs/log"}`)

	beego.Run()
}
