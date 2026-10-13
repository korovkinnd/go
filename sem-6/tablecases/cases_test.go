//go:build !solution

package tablecases

import (
	"errors"
	"testing"
)

func TestShippingCost(t *testing.T) {
	// TODO: составьте таблицу, запустите именованные подтесты и проверьте покрытие.
	for _, tc := range []struct {
		name    string
		weight  int
		express bool
		want    int
		wantErr error
	}{
		{name: "zero weight", weight: 0, wantErr: ErrWeight},
		{name: "negative weight", weight: -5, wantErr: ErrWeight},
		{name: "zero weight express", weight: 0, express: true, wantErr: ErrWeight},
		{name: "min weight", weight: 1, want: 100},
		{name: "min weight express", weight: 1, express: true, want: 250},
		{name: "first tier middle", weight: 3, want: 100},
		{name: "first tier upper bound", weight: 5, want: 100},
		{name: "first tier upper bound express", weight: 5, express: true, want: 250},
		{name: "second tier lower bound", weight: 6, want: 180},
		{name: "second tier lower bound express", weight: 6, express: true, want: 330},
		{name: "second tier upper bound", weight: 10, want: 180},
		{name: "second tier upper bound express", weight: 10, express: true, want: 330},
		{name: "third tier lower bound", weight: 11, want: 300},
		{name: "third tier lower bound express", weight: 11, express: true, want: 450},
		{name: "max weight", weight: 20, want: 300},
		{name: "max weight express", weight: 20, express: true, want: 450},
		{name: "above max weight", weight: 21, wantErr: ErrWeight},
		{name: "above max weight express", weight: 21, express: true, wantErr: ErrWeight},
		{name: "huge weight", weight: 1000, wantErr: ErrWeight},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ShippingCost(tc.weight, tc.express)
			if got != tc.want {
				t.Errorf("ShippingCost(%d, %t) = %d, want %d", tc.weight, tc.express, got, tc.want)
			}

			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Errorf("ShippingCost(%d, %t) error = %v, want %v", tc.weight, tc.express, err, tc.wantErr)
				}
			} else if err != nil {
				t.Errorf("ShippingCost(%d, %t) error = %v, want nil", tc.weight, tc.express, err)
			}
		})
	}
}
