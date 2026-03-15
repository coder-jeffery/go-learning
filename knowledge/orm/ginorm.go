package orm

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name:/:action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		context.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})

	})
	r.Run("8888")
}
