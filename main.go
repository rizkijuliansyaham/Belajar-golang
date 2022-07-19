package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/Hello", helloHandler)

	router.GET("/books/:id", urlParamHandler)

	router.GET("/query", urlQueryHandler)

	router.Run()
	// router.Run(":5655") // perubahan port
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
