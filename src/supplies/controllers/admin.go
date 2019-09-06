package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"supplies/models"
)

type AdminController struct {
	beego.Controller
}

func (a *AdminController) Prepare() {
	a.EnableXSRF = true
}

func (c *AdminController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// @router /index [get]
func (c *AdminController) Index() {
	c.TplName = "admin/index.html"
}

// @route /add [get]
func (c *AdminController) Add() {
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "admin/add.html"
}

// @route /store [post]
func (c *AdminController) Store() {
	var v models.Admin
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = &Response{0, 0, "ok", v}
	}
	c.ServeJSON()
}