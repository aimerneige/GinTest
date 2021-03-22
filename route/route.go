package route

import (
	"aimerneige.com/GinTest/controller"
	"github.com/gin-gonic/gin"
)

func RouteCollect(r *gin.Engine) *gin.Engine {

	r.GET("/", controller.Index)

	r.GET("/book", controller.GetBook)
	r.POST("/book", controller.AddBook)

	return r
}
