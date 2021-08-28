package controllers

import (
	"github.com/SuperDsj/webcron/app/libs"
	"github.com/SuperDsj/webcron/app/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type GatewayController struct {
	BaseController
}

func (this *GatewayController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.TaskGatewayGetList(page, this.pageSize)

	this.Data["pageTitle"] = "分组列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("GatewayController.List"), true).ToString()
	this.display()
}

func (this *GatewayController) Add() {
	if this.isPost() {
		group := new(models.TaskGateway)
		group.GroupName = strings.TrimSpace(this.GetString("group_name"))
		group.InboundType = strings.TrimSpace(this.GetString("inbound_type"))
		group.InboundName = strings.TrimSpace(this.GetString("inbound_name"))
		group.InboundLocal = strings.TrimSpace(this.GetString("inbound_local"))
		group.InboundSip = strings.TrimSpace(this.GetString("inbound_sip"))
		group.InboundFs = strings.TrimSpace(this.GetString("inbound_fs"))
		group.UserId = this.userId
		group.Description = strings.TrimSpace(this.GetString("description"))

		_, err := models.TaskGatewayAdd(group)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加网关"
	this.display()
}

func (this *GatewayController) Edit() {
	id, _ := this.GetInt("id")

	group, err := models.TaskGatewayGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		group.GroupName = strings.TrimSpace(this.GetString("group_name"))
		group.InboundType = strings.TrimSpace(this.GetString("inbound_type"))
		group.InboundName = strings.TrimSpace(this.GetString("inbound_name"))
		group.InboundLocal = strings.TrimSpace(this.GetString("inbound_local"))
		group.InboundSip = strings.TrimSpace(this.GetString("inbound_sip"))
		group.InboundFs = strings.TrimSpace(this.GetString("inbound_fs"))
		group.Description = strings.TrimSpace(this.GetString("description"))
		err := group.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑分组"
	this.Data["group"] = group
	this.display()
}

func (this *GatewayController) Batch() {
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
			models.TaskGatewayDelById(id)
			//models.TaskResetGroupId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
