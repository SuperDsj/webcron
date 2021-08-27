package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type TTaskManger struct {
	Id          	int
	UserId      	int
	MangerName 		string
	InboundType	string
	InboundName	string
	InboundLocal	string
	InboundSip		string
	InboundFs		string
	Description 	string
	CreateTime  	int64
}

type TaskManger struct {
	Id          	int
	UserId      	int
	MangerName 		string
	InboundType	string
	InboundName	string
	InboundLocal	string
	InboundSip		string
	InboundFs		string
	Description 	string
	CreateTime  	int64
}

func (t *TaskManger) MangerTableName() string {
	return TableName("`task_manger`")
}

func (t *TaskManger) MangerUpdate(fields ...string) error {
	if t.MangerName == "" {
		return fmt.Errorf("名不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func TaskMangerAdd(obj *TaskManger) (int64, error) {
	if obj.MangerName == "" {
		return 0, fmt.Errorf("名不能为空")
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
	return obj, nil
}

func TaskMangerDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("task_manger")).Filter("id", id).Delete()
	return err
}

func TaskMangerGetList(page, pageSize int) ([]*TaskManger, int64) {
	offset := (page - 1) * pageSize

	list := make([]*TaskManger, 0)
	fmt.Println(TableName("task_manger"))
	query := orm.NewOrm().QueryTable(TableName("task_manger"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
