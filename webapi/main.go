package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"webapi/page"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	page.Register(r)

	err := r.Run()
	if err != nil {
		os.Exit(1)
	} // listen and serve on 0.0.0.0:8080
}
