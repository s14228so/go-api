package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Golang 入門",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

 