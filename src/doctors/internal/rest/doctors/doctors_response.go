package doctors

import (
	"doctor-search-engine/doctors/internal/domain"
	"net/http"

	"github.com/go-chi/render"
)

type DoctorsResponse struct {
	Doctors []DoctorResponse `json:"doctors"`
}

type DoctorResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	RegistrationId string `json:"registration_id"`
	Phone          string `json:"phone"`
	Adress         string `json:"adress"`
	City           string `json:"city"`
	ZipCode        string `json:"zip_code"`
	Country        string `json:"country"`
	SpecialityId   int    `json:"speciality_id"`
	SpecialityName string `json:"speciality_name"`
}

func NewDoctorResponse(d *domain.Doctor) DoctorResponse {
	return DoctorResponse{
		Id:             d.Id,
		Name:           d.Name,
		Surname:        d.Surname,
		Email:          d.Email,
		RegistrationId: d.RegistrationId,
		Phone:          d.Phone,
		Adress:         d.Adress,
		City:           d.City,
		ZipCode:        d.ZipCode,
		Country:        d.Country,
		SpecialityId:   d.SpecialityId,
		SpecialityName: d.SpecialityName,
	}
}

func NewDoctorsResponse(doctors []*domain.Doctor) DoctorsResponse {
	doctorsResponse := DoctorsResponse{}
	doctorsResponse.Doctors = make([]DoctorResponse, len(doctors))
	for idx, d := range doctors {
		doctorsResponse.Doctors[idx] = NewDoctorResponse(d)
	}

	return doctorsResponse
}

func (DoctorsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}

func (DoctorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}
