package main

import (
	"aimerneige.com/GinTest/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.RouteCollect(r)

	panic(r.Run(":8080"))
}
