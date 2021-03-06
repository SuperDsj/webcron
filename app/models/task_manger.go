package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type TaskManger struct {
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

func (t *TaskManger) TableName() string {
	return TableName("task_manger")
}

func (t *TaskManger) Update(fields ...string) error {
	if t.GroupName == "" {
		return fmt.Errorf("组名不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func TaskMangerAdd(obj *TaskManger) (int64, error) {
	if obj.GroupName == "" {
		return 0, fmt.Errorf("组名不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func TaskMangerGetById(id int) (*TaskManger, error) {
	obj := &TaskManger{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	fmt.Println(obj)
	return obj, nil
}

func TaskMangerDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("task_manger")).Filter("id", id).Delete()
	return err
}

func TaskMangerGetList(page, pageSize int) ([]*TaskManger, int64) {
	offset := (page - 1) * pageSize

	list := make([]*TaskManger, 0)

	query := orm.NewOrm().QueryTable(TableName("task_manger"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
