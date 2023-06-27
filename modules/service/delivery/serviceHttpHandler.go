package delivery

import (
	"log"
	"net/http"
	"strings"

	customError "github.com/gulliverwald/clean-arch-go/error"
	repository "github.com/gulliverwald/clean-arch-go/modules/service/repositories"
	. "github.com/gulliverwald/clean-arch-go/modules/service/usecase"
	"github.com/gulliverwald/clean-arch-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/gulliverwald/clean-arch-go/domain"
)

type ServiceHandler struct {
	ServiceUsecase domain.ServiceUsecase
}

func NewServiceHttpHandler(route *gin.Engine) {
	repository := repository.NewMySqlRepository()
	errorRepository := customError.NewHttpError()
	useCase := NewServiceUsecase(repository, errorRepository)

	handler := &ServiceHandler{
		ServiceUsecase: useCase,
	}

	route.POST("/service", handler.Create)
	route.GET("/service", handler.Fetch)
	route.GET("/service/:id", handler.GetById)
	route.PUT("/service/:id", handler.Update)
	route.DELETE("/service/:id", handler.Delete)
}

func (httpHandler *ServiceHandler) Create(ctx *gin.Context) {
	var service domain.Service

	err := ctx.Bind(&service)
	if err != nil || utils.IsInputValid(&service) {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Input is invalid."})
		return
	}

	err = httpHandler.ServiceUsecase.Create(ctx, &service)
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

	ctx.JSON(http.StatusCreated, service)
}

func (httpHandler *ServiceHandler) GetById(ctx *gin.Context) {
	var request struct {
		ID int64 `uri:"id"`
	}

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid ID."})
		return
	}

	service, err := httpHandler.ServiceUsecase.GetByID(ctx, request.ID)
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

	ctx.JSON(http.StatusOK, service)
}

func (httpHandler *ServiceHandler) Fetch(ctx *gin.Context) {
	services, err := httpHandler.ServiceUsecase.Fetch(ctx)
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

	ctx.JSON(http.StatusOK, services)
}

func (httpHandler *ServiceHandler) Update(ctx *gin.Context) {
	var service domain.Service
	var request struct {
		ID int64 `uri:"id"`
	}

	if err := ctx.ShouldBindUri(&request); err != nil {
		log.Println("#ERROR: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "The was an internal error."})
		return
	}

	ctx.Bind(&service)

	if service.ID != request.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "The service ID and request ID does not match."})
		return
	}

	if utils.IsInputValid(&service) {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Input is invalid."})
		return
	}

	err := httpHandler.ServiceUsecase.Update(ctx, &service)
	if err != nil {
		log.Println("#ERROR: ", err)
		customError, ok := err.(customError.CustomError)

		if ok {
			ctx.JSON(customError.ErrorCode, gin.H{"errorMessage": customError.ErrorMessage})
			return
		}

		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "The service already exists."})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "There was an internal error."})
		return
	}

	ctx.JSON(http.StatusOK, service)
}

func (httpHandler *ServiceHandler) Delete(ctx *gin.Context) {
	var request struct {
		ID int64 `uri:"id"`
	}

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid service ID."})
		return
	}

	err := httpHandler.ServiceUsecase.Delete(ctx, request.ID)
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

	ctx.JSON(http.StatusOK, nil)
}
