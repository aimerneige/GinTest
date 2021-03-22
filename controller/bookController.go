package controller

import (
	"aimerneige.com/GinTest/model"
	"aimerneige.com/GinTest/response"
	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {

	// query data
	name := ctx.Query("name")
	author := ctx.Query("author")
	year := ctx.Query("year")

	newBook := model.Book{
		Name:   name,
		Author: author,
		Year:   year,
	}

	data := gin.H{
		"name":   newBook.Name,
		"author": newBook.Author,
		"year":   newBook.Year,
	}

	message := "Book create successful!"

	response.Success(ctx, data, message)
}

func AddBook(ctx *gin.Context) {

	// form body
	name := ctx.PostForm("name")
	author := ctx.PostForm("author")
	year := ctx.PostForm("year")

	newBook := model.Book{
		Name:   name,
		Author: author,
		Year:   year,
	}

	data := gin.H{
		"name":   newBook.Name,
		"author": newBook.Author,
		"year":   newBook.Year,
	}

	message := "Book create successful!"

	response.Success(ctx, data, message)
}
