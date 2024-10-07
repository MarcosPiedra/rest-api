package queries

import (
	"context"
	"doctor-search-engine/doctors/internal/domain"
)

type (
	GetSpecialitiesHandler struct {
		specialityRepository domain.SpecialityRepository
	}
)

func NewGetSpecialitiesHandler(specialityRepository domain.SpecialityRepository) GetSpecialitiesHandler {
	return GetSpecialitiesHandler{specialityRepository: specialityRepository}
}

func (h GetSpecialitiesHandler) GetSpecialities(ctx context.Context) ([]*domain.Speciality, error) {
	return h.specialityRepository.GetSpecialities(ctx)
}
