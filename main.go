package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "Rizki Juliansyah",
			"bio":  "Belajar Golang dari youtube",
		})
	})
	router.GET("/Hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"content":    "Test path",
			"Keterangan": "Belajar Golang dari youtube tes path",
		})
	})

	// router.Run()
	router.Run(":5655") // perubahan port
}
