package controller

import (
	"aimerneige.com/GinTest/response"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	message := "Server Start Successful"

	response.Success(ctx, nil, message)
}
