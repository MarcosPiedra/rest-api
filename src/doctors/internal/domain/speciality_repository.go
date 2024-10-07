package domain

import (
	"context"
)

type SpecialityRepository interface {
	ExistsSpeciality(id int, ctx context.Context) (bool, error)
	GetSpecialities(ctx context.Context) ([]*Speciality, error)
}
