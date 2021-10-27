package main

import "github.com/gin-gonic/gin"

func main() {
	srv := gin.Default()

	srv.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	srv.Run()
}
