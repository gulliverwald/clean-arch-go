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

func TestCustomerGetByIdUsecase(t *testing.T) {
	mockRepository := repository.NewMockRepository()
	httpError := customError.NewHttpError()
	customerUsecase := NewCustomerUsecase(mockRepository, httpError)

	customer := &domain.Customer{
		Firstname: "Unitary",
		Lastname:  "Test",
		Document:  "SOME_DOCUMENT",
	}

	t.Run("Should return an error on update a customer", func(t *testing.T) {
		want := errors.New("Failed to update customer")
		mockRepository.On("Update").Return(want).Once()

		got := customerUsecase.Update(context.TODO(), customer)

		assert.Equal(t, want, got)
	})

	t.Run("Should update and return a customer successfully", func(t *testing.T) {
		got := customerUsecase.Update(context.TODO(), customer)

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
