package domain

import "context"

type Customer struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Document  string `json:"document"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CustomerUsecase interface {
	Fetch(ctx context.Context) ([]Customer, error)
	GetByID(ctx context.Context, id int64) (Customer, error)
	Create(ctx context.Context, customer *Customer) error
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id int64) error
}

type CustomerRepository interface {
	Fetch(ctx context.Context) ([]Customer, error)
	GetByID(ctx context.Context, id int64) (Customer, error)
	Create(ctx context.Context, customer *Customer) error
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id int64) error
}
