package main

import (
	"net/http"

	handlers "github.com/gulliverwald/clean-arch-go/modules"
	
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.SetTrustedProxies(nil)

	app.Use(handlers.NewCustomerController())

	app.Run()
}