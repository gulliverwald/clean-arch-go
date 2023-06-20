package usecase

import "github.com/gulliverwald/clean-arch-go/domain"

type UseCases struct {
	r domain.CustomerRepository
}

func New(r domain.CustomerRepository) *UseCases {
	return &UseCases{
		repository: r,
	}
}
