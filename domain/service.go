package domain

import "context"

type Service struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Duration     int    `json:"duration"`
	ScheduleDate string `json:"scheduleDate"`
	CustomerID   int    `json:"customerId"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ServiceUsecase interface {
	Fetch(ctx context.Context) ([]Service, error)
	GetByID(ctx context.Context, id int) (Service, error)
	Create(context.Context, *Service) error
	Update(ctx context.Context, ar *Service) error
	Delete(ctx context.Context, id int) error
}

type ServiceRepository interface {
	Fetch(ctx context.Context) ([]Service, error)
	GetByID(ctx context.Context, id int) (Service, error)
	GetServiceByScheduleDateAndCustomerId(ctx context.Context, customerId string, date string) (Service, error)
	Create(ctx context.Context, service *Service) error
	Update(ctx context.Context, service *Service) error
	Delete(ctx context.Context, id int) error
}
