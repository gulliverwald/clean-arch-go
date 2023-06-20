package domain

import "context"

type Service struct {
	ID		  	int64 		`json:"id"`
	Name 	  	string 		`json:"name"`
	Price	  	float32 	`json:"price"`
	CustomerID 	int64		`json:"customerId"`
	CreatedAt 	string 		`json:"created_at"`
	UpdatedAt 	string 		`json:"updated_at"`
}

type ServiceUsecase interface {
	// Fetch(ctx context.Context) ([]Service, string, error)
	// GetByID(ctx context.Context, id int64) (Service, error)
	Create(context.Context, *Service) error
	// Update(ctx context.Context, ar *Service) error
	// Delete(ctx context.Context, id int64) error
}

type ServiceRepository interface {
	// Fetch(ctx context.Context) ([]Service, string, error)
	// GetByID(ctx context.Context, id int64) (Service, error)
	Create(context.Context, *Service) error
	// Update(ctx context.Context, ar *Service) error
	// Delete(ctx context.Context, id int64) error
}