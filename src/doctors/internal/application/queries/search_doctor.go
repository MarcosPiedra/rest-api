package queries

import (
	"context"
	"doctor-search-engine/doctors/internal/domain"
)

type (
	SearchDoctor struct {
		Name         string
		Surname      string
		SpecialityId int
	}
	SearchDoctorHandler struct {
		doctorRepository domain.DoctorRepository
	}
)

func NewSearchDoctorHandler(doctorRepository domain.DoctorRepository) SearchDoctorHandler {
	return SearchDoctorHandler{doctorRepository: doctorRepository}
}

func (h SearchDoctorHandler) SearchDoctors(ctx context.Context, query SearchDoctor) ([]*domain.Doctor, error) {
	return h.doctorRepository.SearchDoctor(ctx, query.Name, query.Surname, query.SpecialityId)
}
