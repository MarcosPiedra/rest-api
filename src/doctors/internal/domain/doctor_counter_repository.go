package domain

import (
	"context"
)

type DoctorCounterRepository interface {
	GetDoctorCounter(ctx context.Context) ([]*DoctorCounter, error)
}
