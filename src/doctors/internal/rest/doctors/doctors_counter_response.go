package doctors

import (
	"doctor-search-engine/doctors/internal/domain"
	"net/http"

	"github.com/go-chi/render"
)

type DoctorsCounterResponse struct {
	Doctors []DoctorCounterResponse `json:"doctors"`
}

type DoctorCounterResponse struct {
	DoctorId     int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	SpecialityId int    `json:"speciality_id"`
}

func NewDoctorCounterResponse(d *domain.DoctorCounter) DoctorCounterResponse {
	return DoctorCounterResponse{
		DoctorId:     d.DoctorId,
		Name:         d.Name,
		Surname:      d.Surname,
		SpecialityId: d.SpecialityId,
	}
}

func NewDoctorsCounterResponse(doctors []*domain.DoctorCounter) DoctorsCounterResponse {
	response := DoctorsCounterResponse{}
	response.Doctors = make([]DoctorCounterResponse, len(doctors))
	for idx, d := range doctors {
		response.Doctors[idx] = NewDoctorCounterResponse(d)
	}

	return response
}

func (DoctorsCounterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}
