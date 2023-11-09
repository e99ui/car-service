package services

import (
	"context"

	"github.com/canter-tech/car-service/internal/domain"
)

type CarRepository interface {
	CreateOrUpdate(ctx context.Context, car *domain.Car) error
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context, id string) (*domain.Car, error)
}

type CarService struct {
	repo CarRepository
}

func NewCarService(repo CarRepository) *CarService {
	return &CarService{
		repo: repo,
	}
}

func (s *CarService) CreateOrUpdate(ctx context.Context, car *domain.Car) error {
	return s.repo.CreateOrUpdate(ctx, car)
}

func (s *CarService) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

func (s *CarService) Get(ctx context.Context, id string) (*domain.Car, error) {
	return s.repo.Get(ctx, id)
}
