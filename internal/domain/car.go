package domain

import "fmt"

type Car struct {
	id       string
	name     string
	class    *string
	brand    string
	yearFrom *int
	yearTo   *int
}

func NewCar(id, name string, class *string, brand string, yearFrom, yearTo *int) (*Car, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: car id is required", ErrRequired)
	}
	if name == "" {
		return nil, fmt.Errorf("%w: car name is required", ErrRequired)
	}
	if brand == "" {
		return nil, fmt.Errorf("%w: car brand is required", ErrRequired)
	}

	return &Car{
		id:       id,
		name:     name,
		class:    class,
		brand:    brand,
		yearFrom: yearFrom,
		yearTo:   yearTo,
	}, nil
}

func (c *Car) ID() string {
	return c.id
}

func (c *Car) Name() string {
	return c.name
}

func (c *Car) Class() *string {
	return c.class
}

func (c *Car) Brand() string {
	return c.brand
}

func (c *Car) YearFrom() *int {
	return c.yearFrom
}

func (c *Car) YearTo() *int {
	return c.yearTo
}
