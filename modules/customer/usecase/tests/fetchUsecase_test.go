package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gulliverwald/clean-arch-go/domain"
	customError "github.com/gulliverwald/clean-arch-go/error"
	repository "github.com/gulliverwald/clean-arch-go/modules/customer/repositories"
	. "github.com/gulliverwald/clean-arch-go/modules/customer/usecase"
)

func TestCustomerFetchUsecase(t *testing.T) {
	mockRepository := repository.NewMockRepository()
	httpError := customError.NewHttpError()
	customerUsecase := NewCustomerUsecase(mockRepository, httpError)

	t.Run("Should fetch all customers successfully", func(t *testing.T) {
		got, gotErr := customerUsecase.Fetch(context.TODO())

		want := []domain.Customer{}

		assert.Equal(t, want, got)
		assert.Equal(t, nil, gotErr)
	})
}
