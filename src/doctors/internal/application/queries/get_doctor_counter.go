package queries

import (
	"context"
	"doctor-search-engine/doctors/internal/domain"
)

type (
	GetDoctorCounterHandler struct {
		doctorSearchCounterRepository domain.DoctorCounterRepository
	}
)

func NewGetDoctorCounterHandler(doctorSearchCounterRepository domain.DoctorCounterRepository) GetDoctorCounterHandler {
	return GetDoctorCounterHandler{doctorSearchCounterRepository: doctorSearchCounterRepository}
}

func (h GetDoctorCounterHandler) GetDoctorsCounter(ctx context.Context) ([]*domain.DoctorCounter, error) {
	return h.doctorSearchCounterRepository.GetDoctorCounter(ctx)
}
