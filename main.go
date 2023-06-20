package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.SetTrustedProxies(nil)

	app.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	app.Run()
}