package main

import (
	_ "auth-demo/config"
	"auth-demo/database"
	_ "auth-demo/database"
	"fmt"
	"github.com/xormplus/core"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, result.Success(gin.H{
	//		"message": "pong",
	//	}))
	//})
	//r.POST("/user/login", handler.Login)
	//auth := r.Group("/")
	//auth.Use(middleware.Auth)
	//{
	//	auth.GET("/secret", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, result.Success("secret"))
	//	})
	//}
	//r.Run(":7777")

	engine:=database.Engine
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	sql:="select.comment.stpl"
	param := map[string]interface{}{"id":"1","comment":"ds"}
	results, _ := engine.SqlTemplateClient(sql, &param).Query().List()
	for _,v := range results {
		fmt.Println(v)
	}
	sql_1_1 := "select * from comment"
	results1, _ := engine.QueryResult(sql_1_1).List()
	for _,v := range results1 {
		fmt.Println(v)
	}
}
