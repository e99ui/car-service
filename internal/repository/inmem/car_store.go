package inmem

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/canter-tech/car-service/internal/domain"
)

type CarStore struct {
	data map[string]*Car
	mu   sync.RWMutex
}

func NewCarStore() *CarStore {
	return &CarStore{
		data: make(map[string]*Car),
	}
}

func (s *CarStore) Get(_ context.Context, id string) (*domain.Car, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	storeCar, exists := s.data[id]
	if !exists {
		return nil, domain.ErrNotFound
	}

	domainCar, err := carStoreToDomain(storeCar)
	if err != nil {
		return nil, fmt.Errorf("portStoreToDomain failed: %w", err)
	}

	return domainCar, nil
}

func (s *CarStore) Count(_ context.Context) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data), nil
}

func (s *CarStore) CreateOrUpdate(ctx context.Context, c *domain.Car) error {
	if c == nil {
		return domain.ErrNil
	}

	storeCar := carDomainToStore(c)

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[storeCar.ID]
	if exists {
		return s.updateCar(ctx, storeCar)
	} else {
		return s.createCar(ctx, storeCar)
	}
}

func (s *CarStore) createCar(_ context.Context, storeCar *Car) error {
	if storeCar == nil {
		return domain.ErrNil
	}

	// set created and updated at
	storeCar.CreatedAt = time.Now()
	storeCar.UpdatedAt = storeCar.CreatedAt

	s.data[storeCar.ID] = storeCar

	return nil
}

func (s *CarStore) updateCar(_ context.Context, c *Car) error {
	if c == nil {
		return domain.ErrNil
	}

	// check if car exists
	storeCar, exists := s.data[c.ID]
	if !exists {
		return domain.ErrNotFound
	}

	storeCar.ID = c.ID
	storeCar.Name = c.Name
	storeCar.Class = c.Class
	storeCar.Brand = c.Brand
	storeCar.YearFrom = c.YearFrom
	storeCar.YearTo = c.YearTo
	storeCar.UpdatedAt = time.Now()

	s.data[c.ID] = storeCar

	return nil
}
