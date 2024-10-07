package domain

import (
	"context"

	"github.com/stackus/errors"
)

var (
	ErrTagNotFound = errors.Wrap(errors.ErrNotFound, "The tag was not found.")
)

type DoctorRepository interface {
	SearchDoctor(ctx context.Context, name string, surname string, specialityId int) ([]*Doctor, error)
	ExistsRegistationId(ctx context.Context, registationId string) (bool, error)
	AddDoctor(ctx context.Context, doctor *Doctor) (int, error)
}
