package custfunc

import "github.com/gin-gonic/gin"

type People struct {
	Name string
	Age  int
}

//
//func HelloWorld(context *gin.Context) {
//
//	s := People{
//		Name: "jeffery",
//		Age:  3,
//	}
//
//	context.HTML(200, "index.html", s) //https://tool.oschina.net/commons?type=5
//	//c.String(200, "Hello World \n")
//}

func Array(context *gin.Context) {
	var arr [3]int
	arr[0] = 99
	arr[1] = 111
	arr[2] = 888
	context.HTML(200, "index.html", arr) //https://tool.oschina.net/commons?type=5
}
