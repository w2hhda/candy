package controllers

import (
	"github.com/w2hhda/candy/models"
	"strings"
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	BaseController
	controllerName string
	actionName     string
	user           *models.Admin
	userId         int
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
}

func (self *AdminController) prepareSideMenu() {
	user := map[string]interface{}{
		"AuthName": "用户管理",
		"Icon":     "fa-user-circle-o",
		"Id":       0,
	}

	userSub := map[string]interface{}{
		"AuthName": "用户管理",
		"Icon":     "fa-user-circle-o",
		"Id":       "17",
		"Pid":      0,
		"AuthUrl":  "/admin/user",
	}

	candy := map[string]interface{}{
		"AuthName": "糖果管理",
		"Icon":     "fa-files-o",
		"Id":       1,
	}

	candySub := map[string]interface{}{
		"AuthName": "糖果类型",
		"Icon":     "fa-files-o",
		"Pid":      1,
		"Id":       "18",
		"AuthUrl":  "/admin/candy",
	}

	trace := map[string]interface{}{
		"AuthName": "财务管理",
		"Icon":     "fa-tree",
		"Id":       2,
	}

	traceSub := map[string]interface{}{
		"AuthName": "财务管理",
		"Icon":     "fa-tree",
		"Pid":      2,
		"Id":       "19",
		"AuthUrl":  "/admin/trace",
	}

	rank := map[string]interface{}{
		"AuthName": "排名管理",
		"Icon":     "fa-code",
		"Id":       3,
	}

	rankSub := map[string]interface{}{
		"AuthName": "排名管理",
		"Icon":     "fa-code",
		"Pid":      3,
		"Id":       "19",
		"AuthUrl":  "/admin/rank",
	}

	self.Data["SideMenu1"] = []map[string]interface{}{
		user, candy, trace, rank,
	}
	self.Data["SideMenu2"] = []map[string]interface{}{
		userSub, candySub, traceSub, rankSub,
	}
}

//前期准备
func (self *AdminController) Prepare() {
	self.pageSize = 20
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["version"] = beego.AppConfig.String("version")
	self.Data["siteName"] = beego.AppConfig.String("site.name")
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName
	fmt.Println(self.controllerName)
	self.Data["loginUserId"] = self.userId
	self.Data["loginUserName"] = self.userName

	self.prepareSideMenu()
}

//加载模板
func (self *AdminController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}

func (self *AdminController) URLMapping() {
	self.Mapping("Index", self.Index)
	self.Mapping("User", self.User)
	self.Mapping("Candy", self.Candy)
	self.Mapping("ListUser", self.ListUser)
	self.Mapping("DisableUser", self.DisableUser)
	self.Mapping("ListCandy", self.ListCandy)
}

// @router /admin/index
func (self *AdminController) Index() {
	self.Data["pageTitle"] = "系统首页"
	self.TplName = "public/main.html"
}

// @router /admin/user
func (self *AdminController) User() {
	self.Layout = "public/layout.html"
	self.TplName = "user/list.html"
}

// @router /admin/candy
func (self *AdminController) Candy() {
	self.Layout = "public/layout.html"
	self.TplName = "candy/list.html"
}

// @router /admin/candy/list [*]
func (self *AdminController) ListCandy() {
	number, _ := strconv.ParseInt(self.Input().Get("page"), 10, 64)
	limit, _ := strconv.ParseInt(self.Input().Get("limit"), 10, 64)

	beego.Info("page=", number)
	// layui page 从1开始, 我们这里从0开始
	page, err := models.ListCandyPage(number-1, limit)
	if err != nil {
		self.RetError(errDB)
	} else {
		self.RetLayuiPage(page.TotalCount, page.List)
	}
}

// @router /admin/user/list [*]
func (self *AdminController) ListUser() {

	number, _ := strconv.ParseInt(self.Input().Get("page"), 10, 64)
	limit, _ := strconv.ParseInt(self.Input().Get("limit"), 10, 64)

	beego.Info("page=", number)
	// layui page 从1开始, 我们这里从0开始
	page, err := models.ListUser(number-1, limit)
	if err != nil {
		self.RetError(errDB)
	} else {
		self.RetLayuiPage(page.TotalCount, page.List)
	}
}

// @router /admin/disable [*]
func (self *AdminController) DisableUser() {
	addr := self.Input().Get("addr")
	status := self.Input().Get("status")
	beego.Info("params", addr, status)

	number, _ := strconv.Atoi(status)
	if number == 0 {
		number = 1
	} else {
		number = 0
	}

	num, err := orm.NewOrm().QueryTable("user").Filter("addr", addr).Update(orm.Params{
		"status": number,
	})
	fmt.Printf("Affected Num: %s, %s", num, err)
	if err == nil {
		self.RetSuccess("")
	} else {
		self.RetError(errDB)
	}
}
