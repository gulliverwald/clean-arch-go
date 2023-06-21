package main

import (
	controller "github.com/gulliverwald/clean-arch-go/modules/customer"
	
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.SetTrustedProxies(nil)

	controller.NewCustomerController(app)

	app.Run()
}
