package controllers

import (
	"github.com/SuperDsj/webcron/app/libs"
	"github.com/SuperDsj/webcron/app/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type MangerController struct {
	BaseController
}

func (this *MangerController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.TaskMangerGetList(page, this.pageSize)

	this.Data["pageTitle"] = "管理列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("MangerController.List"), true).ToString()
	this.display()
}

func (this *MangerController) Add() {
	if this.isPost() {
		group := new(models.TaskManger)
		group.MangerName = strings.TrimSpace(this.GetString("manger_name"))
		group.UserId = this.userId
		group.Description = strings.TrimSpace(this.GetString("description"))

		_, err := models.TaskMangerAdd(group)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加管理"
	this.display()
}

func (this *MangerController) Edit() {
	id, _ := this.GetInt("id")

	group, err := models.TaskMangerGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		group.MangerName = strings.TrimSpace(this.GetString("manger_name"))
		group.Description = strings.TrimSpace(this.GetString("description"))
		err := group.MangerUpdate()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑管理"
	this.Data["group"] = group
	this.display()
}

func (this *MangerController) Batch() {
	action := this.GetString("action")
	ids := this.GetStrings("ids")
	if len(ids) < 1 {
		this.ajaxMsg("请选择要操作的项目", MSG_ERR)
	}

	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		if id < 1 {
			continue
		}
		switch action {
		case "delete":
			models.TaskMangerDelById(id)
			models.TaskResetGroupId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
