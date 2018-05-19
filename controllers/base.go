package controllers

import (
	"github.com/astaxie/beego"
	"github.com/w2hhda/candy/models"
)

type BaseController struct {
	beego.Controller
}

type Data struct {
	models.Request
	Addrs      []string `json:"addrs"`
	PageNumber int64    `json:"page_number"`
}