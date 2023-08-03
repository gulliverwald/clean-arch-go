package usecase

import (
	"context"

	. "github.com/gulliverwald/clean-arch-go/domain"
	customError "github.com/gulliverwald/clean-arch-go/error"
)

type UseCases struct {
	repository   CustomerRepository
	errorHandler customError.ErrorRepository
}

func NewCustomerUsecase(r CustomerRepository, e customError.ErrorRepository) *UseCases {
	return &UseCases{
		repository:   r,
		errorHandler: e,
	}
}

func (ucase *UseCases) Create(ctx context.Context, customer *Customer) error {
	err := ucase.repository.Create(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (ucase *UseCases) Fetch(ctx context.Context) ([]Customer, error) {
	customers, err := ucase.repository.Fetch(ctx)
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (ucase *UseCases) GetByID(ctx context.Context, id int) (Customer, error) {
	customer, err := ucase.repository.GetByID(ctx, id)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (ucase *UseCases) Update(ctx context.Context, customer *Customer) error {
	err := ucase.repository.Update(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (ucase *UseCases) Delete(ctx context.Context, id int) error {
	err := ucase.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
