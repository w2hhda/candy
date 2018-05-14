package main

import (
	"github.com/astaxie/beego"
	_ "github.com/w2hhda/candy/models"
	_ "github.com/w2hhda/candy/controllers"
	_ "github.com/w2hhda/candy/routers"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetLogger("file", `{"filename":"logs/log"}`)

	beego.Run()
}
