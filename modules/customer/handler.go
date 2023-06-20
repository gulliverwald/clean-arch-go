package controller

import (
	"net/http"

	repository "github.com/gulliverwald/clean-arch-go/modules/customer/repository"
	. "github.com/gulliverwald/clean-arch-go/modules/customer/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gulliverwald/clean-arch-go/domain"
)

type CustomerHandler struct {
	CustomerUsecase domain.CustomerUsecase
}

func New(c *gin.Context) {
	useCase = New(&repository.MySqlRepository{})

	handler := &CustomerHandler{
		CustomerUsecase: useCase
	}

	c.POST("/customer", handler.Create)
	c.GET("/customer", handler.Fetch)
}

func (cus *CustomerHandler) Create(c gin.Context) err error {
	var customer domain.Customer

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	err = cus.CustomerUsecase.Create(ctx, &customer)
	
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, article)
}

func (cus *CustomerHandler) Fetch(c gin.Context) err error {
	var customer domain.Customer

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	err = cus.CustomerUsecase.Fetch(ctx, &customer)
	
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}