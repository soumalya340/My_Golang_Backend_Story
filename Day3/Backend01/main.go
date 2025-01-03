package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Be more specific with template loading
	r.LoadHTMLFiles("templates/index.html")
	// Then define routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title":   "Welcome to My Golang Journey",
			"message": "I'm learning Golang and building web applications with Gin framework!",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
