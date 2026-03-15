package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func test(c *gin.Context) {
	c.String(200, "Hello World \n")
	c.JSON(200, gin.H{
		"message": "Hello World this is GinTest",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping is success",
	})
}

// https://ruanyifeng.com/survivor/collapse/where-to-go.html

func main() {
	//url: http://localhost:8080/test
	//对页面渲染效果
	r := gin.Default() //引擎 框架的核心发送机 默认服务器  整个web的服务都是由它驱动
	//r := gin.New()
	r.GET("/test", test)
	r.GET("/ping", ping)
	r.GET("/pong", pong)
	r.GET("/pang", pang)
	r.GET("/user", user)
	//启动引擎 服务器启动  参数和端口
	r.Run(":8080")
}

func user(context *gin.Context) {
	context.JSON(http.StatusOK, User{
		Name: context.Query("name"),
		Age:  18,
	})
}

func pang(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello World this is GinTest",
		"code":    200,
	})
}

func pong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong is success",
	})
}
