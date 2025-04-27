package services

import (
	"gin-server/server/configs/response"
	models "gin-server/server/models/mysql/gorm"

	"github.com/gin-gonic/gin"
)

type userService struct{}

var UserService = userService{}

// 抛出便于单独测试
func (svc userService) TestService(data string) string {
	// 直接使用panic抛出异常，不需要return error，全局异常处理会捕获并返回
	// panic(globals.NewException("other auth service not implemented"))
	return data
}

// ============================================ gin接口处理函数 ============================================

/* 测试接口的处理函数 */
func (svc userService) TestApi(c *gin.Context) {
	// ============================================ 尝鲜测试代码 start ============================================
	// query := c.DefaultQuery("ss", "123")
	// query := c.Query("ss")
	// param := c.Param("test")
	// body := make(map[string]interface{})
	// c.BindJSON(&body)
	// fmt.Println(query, param, body, body["user"])

	// 以下获取参数的方法需要在receive中间件中设置
	// _query, _ := c.Get("query")
	// query, _ := _query.(map[string]interface{})
	// _body, _ := c.Get("body")
	// body := _body.(map[string]interface{})
	// _param, _ := c.Get("param")
	// param := _param.(map[string]interface{})
	// fmt.Println(query["ss"], query, body, param)

	// if true {
	// 	response.Ctx(c).Failed("演示错误返回信息", httpError.NOT_FOUND)
	// 	return
	// }

	// =================================== mongodb ======================================
	// models.Test.InsertOne(models.TestModel{
	// 	Name: "test",
	// })

	// var datas []interface{}
	// test := models.TestModel{
	// 	Name: "test",
	// }
	// for i := 0; i < 1; i++ {
	// 	datas = append(datas, test)
	// }
	// models.Test.InsertMany(datas)

	// response.Ctx(c).Success(models.Test.FindOne(map[string]interface{}{}))
	// response.Ctx(c).Success(models.Test.Find(map[string]interface{}{}))
	// response.Ctx(c).Success(models.Test.Find(
	// 	bson.M{
	// 		"_id": bson.M{
	// 			"$in": []bson.ObjectID{models.TransformId("67b57ff73c3bda009a14b3b3"), models.TransformId("67b57fe8dd94b606b7d4e9f1")},
	// 		},
	// 	},
	// ))
	// models.Test.DeleteOne(map[string]interface{}{"_id": "67b58ca3053d57baa9733047"})

	// =================================== go-sql-driver/mysql ======================================
	// models.Test.QueryMany("select * from test")
	// models.Test.Exec("insert into test (name) values ('test insert')")
	// models.Test.Exec("insert into user (name, mark) values ('lhf1', '2112')")

	// =================================== gorm ======================================
	// filter := &models.TestModel{
	// 	ID: 3,
	// }
	// models.Test.FindMany(filter)
	// models.Test.Paging(filter, 0, 10)
	// response.Ctx(c).Success(models.Test.InsertOne(&models.TestModel{
	// 	Name: "gorm insert test aaa",
	// }))
	models.Test.InsertMany([]*models.TestModel{
		{
			Name: "gorm insert many test123 aasssa",
		},
	})
	// aa := ""
	// print(aa[1])

	// ============================================ 尝鲜测试代码 end ============================================

	// 不需要处理错误，因为svc.TestService中已经直接使用panic抛出错误，全局也有error中间件捕获处理
	response.Ctx(c).Success(svc.TestService("你好"))
}
