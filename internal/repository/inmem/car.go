package inmem

import (
	"errors"
	"time"

	"github.com/canter-tech/car-service/internal/domain"
)

type Car struct {
	ID       string
	Name     string
	Class    *string
	Brand    string
	YearFrom *int
	YearTo   *int

	CreatedAt time.Time
	UpdatedAt time.Time
}

func carStoreToDomain(c *Car) (*domain.Car, error) {
	if c == nil {
		return nil, errors.New("store car is nil")
	}
	return domain.NewCar(c.ID, c.Name, c.Class, c.Brand, c.YearFrom, c.YearTo)
}

func carDomainToStore(c *domain.Car) *Car {
	return &Car{
		ID:       c.ID(),
		Name:     c.Name(),
		Class:    c.Class(),
		Brand:    c.Brand(),
		YearFrom: c.YearFrom(),
		YearTo:   c.YearTo(),
	}
}
