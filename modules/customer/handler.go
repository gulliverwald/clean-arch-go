package controller

import (
	"net/http"
	"log"

	repository "github.com/gulliverwald/clean-arch-go/modules/customer/repositories"
	. "github.com/gulliverwald/clean-arch-go/modules/customer/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gulliverwald/clean-arch-go/domain"
)

type CustomerHandler struct {
	CustomerUsecase domain.CustomerUsecase
}

func NewCustomerController(route *gin.Engine) {
	useCase := NewCustomerUsecase(&repository.MySqlRepository{})

	handler := &CustomerHandler{
		CustomerUsecase: useCase,
	}

	route.POST("/customer", handler.Create)
	route.GET("/customer", handler.Fetch)
	route.GET("/customer/:id", handler.GetById)
	route.PUT("/customer/:id", handler.Update)
	route.DELETE("/customer/:id", handler.Delete)
}

func (cus *CustomerHandler) Create(ctx *gin.Context) {
	var customer domain.Customer
	
	ctx.Bind(&customer)
	err := cus.CustomerUsecase.Create(ctx, &customer)
	
	if err != nil {
		log.Println("#ERROR: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

func (cus *CustomerHandler) GetById(ctx *gin.Context) {
	var request struct {
		ID int64 `uri:"id"`
	}

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	
	customer, err := cus.CustomerUsecase.GetByID(ctx, request.ID)
	
	if err != nil {
		log.Println("#ERROR: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (cus *CustomerHandler) Fetch(ctx *gin.Context) {
	customers, err := cus.CustomerUsecase.Fetch(ctx)
	
	if err != nil {
		log.Println("#ERROR: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (cus *CustomerHandler) Update(ctx *gin.Context) {
	var customer domain.Customer
	var request struct {
		ID int64 `uri:"id"`
	}
	
	if err := ctx.ShouldBindUri(&request); err != nil {
		log.Println("#ERROR: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": err})
		return
	}

	ctx.Bind(&customer)

	if customer.ID != request.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "The customer ID and request ID does not match."})
		return
	}

	err := cus.CustomerUsecase.Update(ctx, &customer)
	
	if err != nil {
		log.Println("#ERROR: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (cus *CustomerHandler) Delete(ctx *gin.Context) {
	var request struct {
		ID int64 `uri:"id"`
	}

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid customer ID."})
		return
	}
	
	err := cus.CustomerUsecase.Delete(ctx, request.ID)
	
	if err != nil {
		log.Println("#ERROR: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}