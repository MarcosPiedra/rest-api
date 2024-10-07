package doctors

import (
	"doctor-search-engine/doctors/internal/domain"
	"net/http"

	"github.com/go-chi/render"
)

type SpecialitiesResponse struct {
	Specialities []SpecialtityResponse `json:"specialities"`
}

type SpecialtityResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewSpecialtityResponse(d *domain.Speciality) SpecialtityResponse {
	return SpecialtityResponse{
		Id:   d.Id,
		Name: d.Name,
	}
}

func NewSpecialitiesResponse(specialities []*domain.Speciality) SpecialitiesResponse {
	response := SpecialitiesResponse{}
	response.Specialities = make([]SpecialtityResponse, len(specialities))
	for idx, s := range specialities {
		response.Specialities[idx] = NewSpecialtityResponse(s)
	}

	return response
}

func (SpecialitiesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}
