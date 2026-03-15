package main

import (
	"github.com/gin-gonic/gin"
	"go_learning/web/custfunc"
)

func main() {
	//gin.New()
	r := gin.Default()
	//加载文件
	//r.LoadHTMLFiles("templates/index.html")
	//加载所有文件
	r.LoadHTMLGlob("web/**/*")
	r.GET("/", custfunc.Array)
	r.Run(":8081")
}
