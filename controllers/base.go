package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}