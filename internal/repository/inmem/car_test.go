package inmem

import (
	"reflect"
	"testing"

	"github.com/canter-tech/car-service/internal/domain"
)

func Test_carStoreToDomain(t *testing.T) {
	type args struct {
		c *Car
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Car
		wantErr bool
	}{
		{
			name: "should return error when car is nil",
			args: args{
				c: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return car domain when car is not nil",
			args: args{
				c: newTestStoreCar(t),
			},
			want:    newTestDomainCar(t),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := carStoreToDomain(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("carStoreToDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("carStoreToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTestStoreCar(t *testing.T) *Car {
	t.Helper()
	testCarString := "test car"
	testCarInt := 2010

	return &Car{
		ID:       testCarString,
		Name:     testCarString,
		Class:    &testCarString,
		Brand:    testCarString,
		YearFrom: &testCarInt,
		YearTo:   &testCarInt,
	}
}

func newTestDomainCar(t *testing.T) *domain.Car {
	t.Helper()
	testCarString := "test car"
	testCarInt := 2010

	car, _ := domain.NewCar(
		testCarString,
		testCarString,
		&testCarString,
		testCarString,
		&testCarInt,
		&testCarInt,
	)

	return car
}
