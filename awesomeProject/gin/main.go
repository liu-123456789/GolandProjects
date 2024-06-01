package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hollegin(ctx *gin.Context) {
	ctx.String(http.StatusOK, "holle,gin")
}

func main() {
	server := gin.Default()

	server.Use(func(context *gin.Context) {
		println("这是第一个middlew")
	}, func(context *gin.Context) {
		println("这是第二个middleware")
	})

	server.GET("/holle", hollegin)

	server.POST("/login", func(context *gin.Context) {
		context.String(http.StatusOK, "hello post")
	})

	//参数路由 通过Param获取
	server.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello"+name)
	})

	//查询参数  通过Query获取
	server.GET("/order", func(context *gin.Context) {
		id := context.Query("id")
		context.String(http.StatusOK, "订单id="+id)
	})

	//通配符路由 通过Param获取（*号不能单独存在）
	server.GET("/views/*.html", func(context *gin.Context) {
		view := context.Param(".html")
		context.String(http.StatusOK, "view是"+view)
	})

	//监听9098端口
	server.Run(":9098")

}
