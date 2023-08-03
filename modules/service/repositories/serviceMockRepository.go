package repository

import (
	"context"

	. "github.com/gulliverwald/clean-arch-go/domain"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
	services []Service
}

func NewMockRepository() ServiceRepository {
	return &MockRepository{
		services: []Service{},
	}
}

func (mock *MockRepository) Create(ctx context.Context, service *Service) error {
	return nil
}

func (mock *MockRepository) Fetch(ctx context.Context) ([]Service, error) {
	return []Service{}, nil
}

func (mock *MockRepository) GetByID(ctx context.Context, id int) (Service, error) {
	return Service{}, nil
}

func (mock *MockRepository) GetServiceByScheduleDateAndCustomerId(
	ctx context.Context,
	customerId string,
	scheduleDate string,
) (Service, error) {
	return Service{}, nil
}

func (mock *MockRepository) Update(ctx context.Context, service *Service) error {
	return nil
}

func (mock *MockRepository) Delete(ctx context.Context, id int) error {
	return nil
}
