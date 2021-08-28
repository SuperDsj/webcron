package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type TaskGateway struct {
	Id          	int
	UserId      	int
	GroupName   	string
	InboundType		string
	InboundName		string
	InboundLocal	string
	InboundSip		string
	InboundFs		string
	Description 	string
	CreateTime  	int64
}

func (t *TaskGateway) TableName() string {
	return TableName("task_gateway")
}

func (t *TaskGateway) Update(fields ...string) error {
	if t.GroupName == "" {
		return fmt.Errorf("组名不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func TaskGatewayAdd(obj *TaskGateway) (int64, error) {
	if obj.GroupName == "" {
		return 0, fmt.Errorf("组名不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func TaskGatewayGetById(id int) (*TaskGateway, error) {
	obj := &TaskGateway{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	fmt.Println(obj)
	return obj, nil
}

func TaskGatewayDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("task_gateway")).Filter("id", id).Delete()
	return err
}

func TaskGatewayGetList(page, pageSize int) ([]*TaskGateway, int64) {
	offset := (page - 1) * pageSize

	list := make([]*TaskGateway, 0)

	query := orm.NewOrm().QueryTable(TableName("task_gateway"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
