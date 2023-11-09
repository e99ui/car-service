package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/canter-tech/car-service/internal/domain"
	"github.com/canter-tech/car-service/internal/transport/http/v1/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CarService interface {
	CreateOrUpdate(ctx context.Context, car *domain.Car) error
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context, id string) (*domain.Car, error)
}

type carRoutes struct {
	service CarService
}

func newCarRoutes(service CarService) chi.Router {
	routes := &carRoutes{
		service: service,
	}

	handler := chi.NewRouter()
	handler.Get("/{id}", routes.get)
	handler.Get("/count", routes.count)
	handler.Post("/upload", routes.upload)

	return handler
}

type response struct{}

func (res *response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CarResponse struct {
	Car models.Car
	response
}

func (routes *carRoutes) get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	car, err := routes.service.Get(r.Context(), idParam)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			_ = render.Render(w, r, ErrNotFound)
			return
		}

		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	_ = render.Render(w, r, &CarResponse{
		Car: models.Car{
			ID:       car.ID(),
			Name:     car.Name(),
			Class:    car.Class(),
			Brand:    car.Brand(),
			YearFrom: car.YearFrom(),
			YearTo:   car.YearTo(),
		},
	})
}

type CountResponse struct {
	response
	Count int `json:"count"`
}

func (routes *carRoutes) count(w http.ResponseWriter, r *http.Request) {
	count, err := routes.service.Count(r.Context())
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	_ = render.Render(w, r, &CountResponse{Count: count})
}

func (routes *carRoutes) upload(w http.ResponseWriter, r *http.Request) {
	carChan := make(chan models.Car)
	errChan := make(chan error)
	doneChan := make(chan struct{})

	go func() {
		err := readCars(r.Context(), r.Body, carChan)
		if err != nil {
			errChan <- err
		} else {
			doneChan <- struct{}{}
		}
	}()

	carCount := 0

	for {
		select {
		case <-r.Context().Done():
			return
		case <-doneChan:
			_ = render.Render(w, r, &CountResponse{Count: carCount})
			return
		case err := <-errChan:
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		case car := <-carChan:
			carCount++

			c, err := car.ToDomain()
			if err != nil {
				_ = render.Render(w, r, ErrInvalidRequest(err))
				return
			}

			if err := routes.service.CreateOrUpdate(r.Context(), c); err != nil {
				_ = render.Render(w, r, ErrInvalidRequest(err))
				return
			}
		}
	}
}

type carJson struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	CyrillicName string       `json:"cyrillic-name"`
	Popular      bool         `json:"popular"`
	Country      string       `json:"country"`
	Models       []models.Car `json:"models"`
}

func readCars(ctx context.Context, r io.Reader, carChan chan models.Car) error {
	decoder := json.NewDecoder(r)

	// Read opening delimitr
	t, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening delimitr: %w", err)
	}

	if t != json.Delim('[') {
		return fmt.Errorf("expected [, got %v", err)
	}

	for decoder.More() {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		var carJson carJson
		if err := decoder.Decode(&carJson); err != nil {
			return fmt.Errorf("failed to decode car: %w", err)
		}

		for _, car := range carJson.Models {
			car.Brand = carJson.Name
			carChan <- car
		}
	}

	return nil
}
