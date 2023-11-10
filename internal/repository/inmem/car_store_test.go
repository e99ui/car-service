package inmem

import (
	"context"
	"testing"

	"github.com/canter-tech/car-service/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCarStore_CreateOrUpdate(t *testing.T) {
	store := NewCarStore()

	t.Run("create car", func(t *testing.T) {
		t.Parallel()

		randomCar := newRandomDomainCar(t)
		err := store.CreateOrUpdate(context.Background(), randomCar)
		require.NoError(t, err)

		car, err := store.Get(context.Background(), randomCar.ID())
		require.NoError(t, err)

		require.Equal(t, car, randomCar)
	})

	t.Run("update car", func(t *testing.T) {
		t.Parallel()

		randomCar := newRandomDomainCar(t)
		err := store.CreateOrUpdate(context.Background(), randomCar)
		require.NoError(t, err)

		car, err := store.Get(context.Background(), randomCar.ID())
		require.NoError(t, err)

		require.Equal(t, car, randomCar)

		beforeUpdateCar := newUpdatedDomainCar(t, car.ID())

		err = store.CreateOrUpdate(context.Background(), beforeUpdateCar)
		require.NoError(t, err)

		updatedCar, err := store.Get(context.Background(), randomCar.ID())
		require.NoError(t, err)

		require.Equal(t, randomCar.ID(), updatedCar.ID())
		require.NotEqual(t, randomCar.Name(), updatedCar.Name())
	})

	t.Run("nil car", func(t *testing.T) {
		t.Parallel()
		err := store.CreateOrUpdate(context.Background(), nil)
		require.ErrorIs(t, err, domain.ErrNil)
	})
}

func newRandomDomainCar(t *testing.T) *domain.Car {
	t.Helper()
	testID := uuid.New()
	testStr := "test car"
	testInt := 2010
	car, _ := domain.NewCar(testID.String(), testStr, &testStr, testStr, &testInt, &testInt)

	return car
}

func newUpdatedDomainCar(t *testing.T, carID string) *domain.Car {
	t.Helper()
	testStr := "test updated car"
	testInt := 2012
	car, _ := domain.NewCar(carID, testStr, &testStr, testStr, &testInt, &testInt)

	return car

}
