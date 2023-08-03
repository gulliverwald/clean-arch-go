package repository

import (
	"context"
	"errors"
	"time"

	. "github.com/gulliverwald/clean-arch-go/domain"
	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/slices"
)

type MockRepository struct {
	mock.Mock
	customers []Customer
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		customers: []Customer{},
	}
}

func (mock *MockRepository) Create(ctx context.Context, customer *Customer) error {
	newCustomer := *customer

	newCustomer.ID = len(mock.customers)
	newCustomer.CreatedAt = time.Now().Format(time.RFC3339)
	newCustomer.UpdatedAt = time.Now().Format(time.RFC3339)

	mock.customers = append(mock.customers, *customer)

	return nil
}

func (mock *MockRepository) Fetch(ctx context.Context) ([]Customer, error) {
	return mock.customers, nil
}

func (mock *MockRepository) GetByID(ctx context.Context, id int) (Customer, error) {
	index := slices.IndexFunc(mock.customers, func(mCustomer Customer) bool {
		return mCustomer.ID == id
	})

	if index == -1 {
		return Customer{}, errors.New("Customer not found")
	}

	customer := mock.customers[index]

	return customer, nil
}

func (mock *MockRepository) Update(ctx context.Context, customer *Customer) error {
	index := slices.IndexFunc(mock.customers, func(mCustomer Customer) bool {
		return mCustomer.ID == customer.ID
	})

	if index == -1 {
		return errors.New("Customer not found")
	}

	customer.UpdatedAt = time.Now().Format(time.RFC3339)

	mock.customers[index] = *customer

	return nil
}

func (mock *MockRepository) Delete(ctx context.Context, id int) error {
	index := slices.IndexFunc(mock.customers, func(mCustomer Customer) bool {
		return mCustomer.ID == id
	})

	if index == -1 {
		return errors.New("Customer not found")
	}

	mock.customers = slices.Delete(mock.customers, index, 1)

	return nil
}
