package usecase

import (
	"context"
	"fmt"

	. "github.com/gulliverwald/clean-arch-go/domain"
)

type UseCases struct {
	repository CustomerRepository
}

func NewCustomerUsecase(r CustomerRepository) *UseCases {
	return &UseCases{
		repository: r,
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
	fmt.Printf("buscar\n")

	return []Customer{}, nil
}

func (ucase *UseCases) GetByID(ctx context.Context, id int64) (Customer, error) {
	fmt.Printf("busquei pelo id\n")
	fmt.Printf("id: %d\n", id)

	return Customer{}, nil
}

func (ucase *UseCases) Update(context.Context, *Customer) error {
	fmt.Printf("atualizar\n")

	return nil
}

func (ucase *UseCases) Delete(ctx context.Context, id int64) error {
	fmt.Printf("deletar\n")

	return nil
}
