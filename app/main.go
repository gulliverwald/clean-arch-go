package main

import (
	customerDelivery "github.com/gulliverwald/clean-arch-go/modules/customer/delivery"
	serviceDelivery "github.com/gulliverwald/clean-arch-go/modules/service/delivery"

	"github.com/gin-gonic/gin"
)

func initializeRESTful() {
	app := gin.Default()
	app.SetTrustedProxies(nil)

	customerDelivery.NewCustomerHttpHandler(app)
	serviceDelivery.NewServiceHttpHandler(app)

	app.Run()
}

func main() {
	initializeRESTful()
}
