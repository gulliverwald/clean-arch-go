package usecase

import (
	"context"
	"net/http"

	. "github.com/gulliverwald/clean-arch-go/domain"
	customError "github.com/gulliverwald/clean-arch-go/error"
	"github.com/rickb777/date"
)

type UseCases struct {
	repository      ServiceRepository
	errorRepository customError.ErrorRepository
}

func NewServiceUsecase(
	r ServiceRepository,
	er customError.ErrorRepository,
) *UseCases {
	return &UseCases{
		repository:      r,
		errorRepository: er,
	}
}

func (ucase *UseCases) Create(ctx context.Context, service *Service) error {
	parsedDate, err := date.ParseISO(service.ScheduleDate)
	if err != nil {
		return err
	}

	today := date.Today()

	if parsedDate.Before(today) {
		return *ucase.errorRepository.New(
			http.StatusBadRequest,
			"Invalid schedule date.",
		)
	}

	if service.Duration > 120 || service.Duration <= 0 {
		return *ucase.errorRepository.New(
			http.StatusBadRequest,
			"The duration of a service must be between 1 and 120 minutes.",
		)
	}

	err = ucase.repository.Create(ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func (ucase *UseCases) Fetch(ctx context.Context) ([]Service, error) {
	services, err := ucase.repository.Fetch(ctx)
	if err != nil {
		return services, err
	}

	return services, nil
}

func (ucase *UseCases) GetByID(ctx context.Context, id int) (Service, error) {
	service, err := ucase.repository.GetByID(ctx, id)
	if err != nil {
		return service, err
	}

	return service, nil
}

func (ucase *UseCases) Update(ctx context.Context, service *Service) error {
	parsedDate, err := date.ParseISO(service.ScheduleDate)
	if err != nil {
		return err
	}

	today := date.Today()

	if parsedDate.Before(today) {
		return *ucase.errorRepository.New(
			http.StatusBadRequest,
			"Invalid schedule date.",
		)
	}

	if service.Duration > 120 || service.Duration <= 0 {
		return *ucase.errorRepository.New(
			http.StatusBadRequest,
			"The duration of a service must be between 1 and 120 minutes.",
		)
	}

	err = ucase.repository.Update(ctx, service)
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
