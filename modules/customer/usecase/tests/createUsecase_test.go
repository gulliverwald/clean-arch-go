package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	domain "github.com/gulliverwald/clean-arch-go/domain"
	customError "github.com/gulliverwald/clean-arch-go/error"
	repository "github.com/gulliverwald/clean-arch-go/modules/customer/repositories"
	. "github.com/gulliverwald/clean-arch-go/modules/customer/usecase"
)

func TestCustomerCreateUsecase(t *testing.T) {
	mockRepository := repository.NewMockRepository()
	httpError := customError.NewHttpError()
	customerUsecase := NewCustomerUsecase(mockRepository, httpError)

	customer := &domain.Customer{
		Firstname: "Unitary",
		Lastname:  "Test",
		Document:  "SOME_DOCUMENT",
	}

	t.Run("Should return an error on create a customer", func(t *testing.T) {
		want := errors.New("Failed to insert customer")
		mockRepository.On("Create", context.TODO(), customer).Return(want).Once()

		got := customerUsecase.Create(context.TODO(), customer)

		assert.Equal(t, want, got)
	})

	t.Run("Should create and return a customer successfully", func(t *testing.T) {
		got := customerUsecase.Create(context.TODO(), customer)

		want := &domain.Customer{
			ID:        1,
			Firstname: "Unitary",
			Lastname:  "Test",
			Document:  "SOME_DOCUMENT",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}

		assert.Equal(t, want, got)
	})
}
