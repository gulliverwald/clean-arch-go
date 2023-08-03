package delivery

import (
	"log"
	"net/http"
	"strings"

	customError "github.com/gulliverwald/clean-arch-go/error"
	repository "github.com/gulliverwald/clean-arch-go/modules/customer/repositories"
	. "github.com/gulliverwald/clean-arch-go/modules/customer/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gulliverwald/clean-arch-go/domain"
)

type CustomerHandler struct {
	CustomerUsecase domain.CustomerUsecase
}

func NewCustomerHttpHandler(route *gin.Engine) {
	repository := repository.NewMySqlRepository()
	errorRepository := customError.NewHttpError()
	useCase := NewCustomerUsecase(repository, errorRepository)

	handler := &CustomerHandler{
		CustomerUsecase: useCase,
	}

	route.POST("/customer", handler.Create)
	route.GET("/customer", handler.Fetch)
	route.GET("/customer/:id", handler.GetById)
	route.PUT("/customer/:id", handler.Update)
	route.DELETE("/customer/:id", handler.Delete)
}

func (httpHandler *CustomerHandler) Create(ctx *gin.Context) {
	var customer domain.Customer

	ctx.Bind(&customer)
	err := httpHandler.CustomerUsecase.Create(ctx, &customer)
	if err != nil {
		log.Println("#ERROR: ", err)
		customError, ok := err.(customError.CustomError)

		if ok {
			ctx.JSON(customError.ErrorCode, gin.H{"errorMessage": customError.ErrorMessage})
			return
		}

		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Customer with document already exists."})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

func (httpHandler *CustomerHandler) GetById(ctx *gin.Context) {
	var request struct {
		ID int `uri:"id"`
	}

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid input."})
		return
	}

	customer, err := httpHandler.CustomerUsecase.GetByID(ctx, request.ID)
	if err != nil {
		log.Println("#ERROR: ", err)
		customError, ok := err.(customError.CustomError)

		if ok {
			ctx.JSON(customError.ErrorCode, gin.H{"errorMessage": customError.ErrorMessage})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (httpHandler *CustomerHandler) Fetch(ctx *gin.Context) {
	customers, err := httpHandler.CustomerUsecase.Fetch(ctx)
	if err != nil {
		log.Println("#ERROR: ", err)
		customError, ok := err.(customError.CustomError)

		if ok {
			ctx.JSON(customError.ErrorCode, gin.H{"errorMessage": customError.ErrorMessage})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (httpHandler *CustomerHandler) Update(ctx *gin.Context) {
	var customer domain.Customer
	var request struct {
		ID int `uri:"id"`
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

	err := httpHandler.CustomerUsecase.Update(ctx, &customer)
	if err != nil {
		log.Println("#ERROR: ", err)
		customError, ok := err.(customError.CustomError)

		if ok {
			ctx.JSON(customError.ErrorCode, gin.H{"errorMessage": customError.ErrorMessage})
			return
		}

		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Customer with document already exists."})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (httpHandler *CustomerHandler) Delete(ctx *gin.Context) {
	var request struct {
		ID int `uri:"id"`
	}

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid customer ID."})
		return
	}

	err := httpHandler.CustomerUsecase.Delete(ctx, request.ID)
	if err != nil {
		log.Println("#ERROR: ", err)
		customError, ok := err.(customError.CustomError)

		if ok {
			ctx.JSON(customError.ErrorCode, gin.H{"errorMessage": customError.ErrorMessage})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}
