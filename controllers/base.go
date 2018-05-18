package controllers

import (
	"github.com/astaxie/beego"
	"github.com/w2hhda/candy/models"
	"strconv"
)

type BaseController struct {
	beego.Controller
}

type Data struct {
	models.Request
	Addrs      []string `json:"addrs"`
	PageNumber int64    `json:"page_number"`
}

func parseString(input float64) string {
	if input == 0 {
		return "0"
	}
	return strconv.FormatFloat(input, 'E', -1, 64)
}

func parseFloat(input string) float64 {
	out, _ := strconv.ParseFloat(input, 64)
	return out
}