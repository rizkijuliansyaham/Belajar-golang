package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/Hello", helloHandler)
	router.GET("/books/:id", urlParamHandler)
	router.GET("/query", urlQueryHandler)

	router.POST("/books", postBookHandler)

	// router.Run()
	router.Run(":5655") // perubahan port
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	// SubTitle string `json:"sub_title"`
}

func postBookHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		// log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		// "sub_title": bookInput.SubTitle,
	})

}

func urlQueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func urlParamHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Rizki Juliansyah",
		"bio":  "Belajar Golang dari youtube",
	})

}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"content":    "Test path",
		"Keterangan": "Belajar Golang dari youtube tes path",
	})

}
