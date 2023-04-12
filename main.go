package main

import "github.com/gin-gonic/gin"

func main() {
	// print hello world
	println("Hello World")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
