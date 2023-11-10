package domain

import (
	"reflect"
	"testing"
)

func TestNewCar(t *testing.T) {
	type args struct {
		id       string
		name     string
		class    *string
		brand    string
		yearFrom *int
		yearTo   *int
	}

	carID := "test"
	carName := "test Name"
	carClass := "test Class"
	carBrand := "test Brand"
	carYearFrom := 2010
	carYearTo := 2020

	tests := []struct {
		name    string
		args    args
		want    *Car
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid",
			args: args{
				id:       carID,
				name:     carName,
				class:    &carClass,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   &carYearTo,
			},
			want: &Car{
				id:       carID,
				name:     carName,
				class:    &carClass,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   &carYearTo,
			},
			wantErr: false,
		},
		{
			name: "missing car id",
			args: args{
				id:       "",
				name:     carName,
				class:    &carClass,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   &carYearTo,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing car name",
			args: args{
				id:       carName,
				name:     "",
				class:    &carClass,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   &carYearTo,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing car class",
			args: args{
				id:       carID,
				name:     carName,
				class:    nil,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   &carYearTo,
			},
			want: &Car{
				id:       carID,
				name:     carName,
				class:    nil,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   &carYearTo,
			},
			wantErr: false,
		},
		{
			name: "missing car brand",
			args: args{
				id:       carID,
				name:     carName,
				class:    &carClass,
				brand:    "",
				yearFrom: &carYearFrom,
				yearTo:   &carYearTo,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing car year from",
			args: args{
				id:       carID,
				name:     carName,
				class:    &carClass,
				brand:    carBrand,
				yearFrom: nil,
				yearTo:   &carYearTo,
			},
			want: &Car{
				id:       carID,
				name:     carName,
				class:    &carClass,
				brand:    carBrand,
				yearFrom: nil,
				yearTo:   &carYearTo,
			},
			wantErr: false,
		},
		{
			name: "missing car year to",
			args: args{
				id:       carID,
				name:     carName,
				class:    &carClass,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   nil,
			},
			want: &Car{
				id:       carID,
				name:     carName,
				class:    &carClass,
				brand:    carBrand,
				yearFrom: &carYearFrom,
				yearTo:   nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewCar(tt.args.id, tt.args.name, tt.args.class, tt.args.brand, tt.args.yearFrom, tt.args.yearTo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCar() = %v, want %v", got, tt.want)
			}
		})
	}
}
