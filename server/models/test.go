package models

import (
	dbs "gin-server/server/dbs"
)

type test struct {
	BaseModel
}

type TestModel struct {
	Name string `json:"name"`
}

var Test test

func (t *test) InitModel() {
	Test = test{
		BaseModel: BaseModel{
			Colllection: dbs.Mongodb.Collection("test_table"),
		},
	}
}
