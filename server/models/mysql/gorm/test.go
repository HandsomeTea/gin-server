package models

import "gin-server/server/dbs"

type test struct {
	BaseModel
}

type TestModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Test test

func (t *test) InitModel() {
	Test = test{
		BaseModel: BaseModel{
			table:          dbs.Mysql.Table("test"),
			tableInterface: &TestModel{},
		},
	}
}
