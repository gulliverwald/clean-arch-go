package main

import (
	customerHttp "github.com/gulliverwald/clean-arch-go/modules/customer/delivery"
	serviceHttp "github.com/gulliverwald/clean-arch-go/modules/service/delivery"

	"github.com/gin-gonic/gin"
)

func initializeRESTful() {
	app := gin.Default()
	app.SetTrustedProxies(nil)

	customerHttp.NewCustomerHttpHandler(app)
	serviceHttp.NewServiceHttpHandler(app)

	app.Run()
}

func main() {
	initializeRESTful()
}
