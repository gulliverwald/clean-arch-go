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

func TestCustomerUpdateUsecase(t *testing.T) {
	mockRepository := repository.NewMockRepository()
	httpError := customError.NewHttpError()
	customerUsecase := NewCustomerUsecase(mockRepository, httpError)

	customer := &domain.Customer{
		Firstname: "Unitary",
		Lastname:  "Test",
		Document:  "SOME_DOCUMENT",
	}

	t.Run("Should return an error on update a customer", func(t *testing.T) {
		want := errors.New("Customer not found")
		mockRepository.On("Update").Return(want).Once()

		got := customerUsecase.Update(context.TODO(), customer)

		assert.Equal(t, want, got)
	})

	t.Run("Should update and return a customer successfully", func(t *testing.T) {
		mockRepository.Create(context.TODO(), customer)

		updatedCustomer := &domain.Customer{
			ID:        1,
			Firstname: "Unitary",
			Lastname:  "Tester",
			Document:  "SOME_DOCUMENT",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}

		got := customerUsecase.Update(context.TODO(), updatedCustomer)

		assert.Equal(t, nil, got)
	})
}
