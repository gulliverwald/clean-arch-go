package domain

import "context"

type Customer struct {
	ID		  int64 	`json:"id"`
	Name 	  string 	`json:"name"`
	CreatedAt string 	`json:"created_at"`
	UpdatedAt string 	`json:"updated_at"`
}

type CustomerUsecase interface {
	// Fetch(ctx context.Context) ([]Service, string, error)
	// GetByID(ctx context.Context, id int64) (Service, error)
	Create(context.Context, *Service) error
	// Update(ctx context.Context, ar *Service) error
	// Delete(ctx context.Context, id int64) error
}

type CustomerRepository interface {
	// Fetch(ctx context.Context) ([]Service, string, error)
	// GetByID(ctx context.Context, id int64) (Service, error)
	Create(context.Context, *Service) error
	// Update(ctx context.Context, ar *Service) error
	// Delete(ctx context.Context, id int64) error
}