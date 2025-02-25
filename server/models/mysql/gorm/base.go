package models

import "gorm.io/gorm"

type BaseModel struct {
	table          *gorm.DB
	tableInterface interface{}
}

/* 仅支持数据库自带的自增id */
func (b *BaseModel) FindById(id any) map[string]interface{} {
	var result map[string]interface{}

	b.table.First(b.tableInterface, id).Scan(&result)

	return result
}

func (b *BaseModel) FindMany(filter interface{}) []map[string]interface{} {
	var results []map[string]interface{}

	b.table.Where(filter).Find(&results)
	return results
}

func (b *BaseModel) FindOne(filter interface{}) map[string]interface{} {
	var result map[string]interface{}

	b.table.Where(filter).First(b.tableInterface).Scan(&result)

	return result
}

func (b *BaseModel) Paging(filter interface{}, skip int, limit int) interface{} {
	var list []map[string]interface{}
	var count int64

	b.table.Where(filter).Offset(skip).Limit(limit).Find(&list)
	b.table.Where(filter).Model(b.tableInterface).Count(&count)
	type pageResult struct {
		Total int64                    `json:"total"`
		List  []map[string]interface{} `json:"list"`
	}
	result := pageResult{
		Total: count,
		List:  list,
	}
	return result
}

func (b *BaseModel) InsertOne(data interface{}) map[string]interface{} {
	var result map[string]interface{}
	b.table.Create(data).Scan(&result)

	return result
}

func (b *BaseModel) InsertMany(data interface{}) {
	b.table.Create(data)
}
