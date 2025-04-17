package models

// import (
// 	"gin-server/server/dbs"
// 	"reflect"
// )

// type BaseModel struct {
// 	model interface{}
// }

// func dealModelField(model interface{}) ([]interface{}, reflect.Value) {
// 	elem := reflect.New(reflect.TypeOf(model).Elem()).Elem()
// 	fields := reflect.TypeOf(model).Elem().NumField()
// 	dest := make([]interface{}, fields)

// 	for i := 0; i < fields; i++ {
// 		dest[i] = elem.Field(i).Addr().Interface()
// 	}

// 	return dest, elem
// }

// func (b *BaseModel) QueryMany(selectSql string) []interface{} {
// 	rows, _ := dbs.Mysql.Query(selectSql)
// 	dest, inter := dealModelField(b.model)
// 	var results []interface{}

// 	for rows.Next() {
// 		rows.Scan(dest...)
// 		results = append(results, inter.Interface())
// 	}
// 	return results
// }

// func (b *BaseModel) QueryOne(selectSql string) interface{} {
// 	row := dbs.Mysql.QueryRow(selectSql)
// 	dest, inter := dealModelField(b.model)

// 	row.Scan(dest...)

// 	return inter.Interface()
// }

// func (b *BaseModel) Exec(sql string) {
// 	dbs.Mysql.Exec(sql)
// }
