package models

import "github.com/canter-tech/car-service/internal/domain"

type Car struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Class    *string `json:"class,omitempty"`
	Brand    string  `json:"brand"`
	YearFrom *int    `json:"year-from"`
	YearTo   *int    `json:"year-to"`
}

func (c *Car) ToDomain() (*domain.Car, error) {
	return domain.NewCar(c.ID, c.Name, c.Class, c.Brand, c.YearFrom, c.YearTo)
}
