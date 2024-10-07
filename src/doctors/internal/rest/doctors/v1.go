package doctors

import (
	"doctor-search-engine/doctors/internal/application"
	"doctor-search-engine/doctors/internal/application/commands"
	"doctor-search-engine/doctors/internal/application/queries"
	we "doctor-search-engine/internal/web"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/stackus/errors"
)

type DoctorsV1 struct {
	app application.App
}

func NewDoctorsV1(app application.App) *DoctorsV1 {
	return &DoctorsV1{
		app: app,
	}
}

func (c *DoctorsV1) Register(r chi.Router) {
	r.Get("/doctors", c.searchDoctor)
	r.Post("/doctors", c.addDoctor)
	r.Get("/doctors/counter", c.getDoctorSearchCounter)
	r.Get("/specialities", c.getSpecialities)
}

// ListImageLocalizations godoc
// @Summary      Search doctor
// @Description  Search doctor
// @Tags         Doctors
// @Produce      json
// @Param        name          query    string     false    "Name of the doctor."
// @Param        surname       query    string     false    "Surname of the doctor."
// @Param        specialityId  query    int        false    "Speciality of the doctor"
// @Success      200  {object}  doctors.DoctorsResponse
// @Failure      500  {object}  we.ErrorResponse
// @Router       /doctors/v1/doctors [get]
func (h *DoctorsV1) searchDoctor(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	specialityId, _ := strconv.Atoi(r.URL.Query().Get("specialityId"))

	ctx := r.Context()

	q := queries.SearchDoctor{
		Name:         name,
		Surname:      surname,
		SpecialityId: specialityId,
	}

	doctors, err := h.app.SearchDoctors(ctx, q)

	if err == nil {
		err = render.Render(w, r, NewDoctorsResponse(doctors))
	}

	if err != nil {
		render.Render(w, r, we.Err(
			errors.ErrInternalServerError.Msg("Internal server error"),
			errors.ErrInternalServerError.HTTPCode()))
	}
}

// addDoctor godoc
// @Summary      Add doctor
// @Description  Add doctor
// @Tags         Doctors
// @Produce      json
// @Param        request    body    doctors.DoctorRequest    true    "Doctor to add"
// @Success      200  {object}  doctors.DoctorsResponse
// @Failure      400  {object}  we.ErrorResponse
// @Failure      500  {object}  we.ErrorResponse
// @Failure      409  {object}  we.ErrorResponse
// @Router       /doctors/v1/doctors [post]
func (h *DoctorsV1) addDoctor(w http.ResponseWriter, r *http.Request) {
	var doctorRequest DoctorRequest

	err := json.NewDecoder(r.Body).Decode(&doctorRequest)
	if err != nil {
		render.Render(w, r, we.Err(
			errors.ErrInternalServerError.Msg("Internal server error"),
			errors.ErrInternalServerError.HTTPCode()))
		return
	}

	ctx := r.Context()

	var addDoctorCmd = commands.AddDoctor{
		Name:           doctorRequest.Name,
		Surname:        doctorRequest.Surname,
		Email:          doctorRequest.Email,
		RegistrationId: doctorRequest.RegistrationId,
		Phone:          doctorRequest.Phone,
		Adress:         doctorRequest.Adress,
		City:           doctorRequest.City,
		ZipCode:        doctorRequest.ZipCode,
		Country:        doctorRequest.Country,
		SpecialityId:   doctorRequest.SpecialityId,
	}

	v := validator.New()
	err = v.Struct(addDoctorCmd)
	if err != nil {
		render.Render(w, r, we.Err(
			errors.ErrBadRequest.Err(err),
			errors.ErrBadRequest.HTTPCode()))

		return
	}

	doctorAdded, err := h.app.AddDoctor(ctx, addDoctorCmd)

	if err == nil {
		render.Render(w, r, NewDoctorResponse(doctorAdded))
	} else {
		var coder errors.HTTPCoder = nil
		if errors.As(err, &coder) {
			render.Render(w, r, we.Err(err, coder.HTTPCode()))
		} else {
			render.Render(w, r, we.Err(
				errors.ErrInternalServerError.Msg("Internal server error"),
				errors.ErrInternalServerError.HTTPCode()))
		}
	}
}

// getDoctorCounter godoc
// @Summary      Get doctor counter
// @Description  Get doctor counter
// @Tags         Doctors
// @Produce      json
// @Success      200  {object}  doctors.DoctorsCounterResponse
// @Failure      500  {object}  we.ErrorResponse
// @Router       /doctors/v1/doctors/counter [get]
func (h *DoctorsV1) getDoctorSearchCounter(w http.ResponseWriter, r *http.Request) {

	doctorsCounter, err := h.app.GetDoctorsCounter(r.Context())

	if err == nil {
		render.Render(w, r, NewDoctorsCounterResponse(doctorsCounter))
	} else {
		render.Render(w, r, we.Err(
			errors.ErrInternalServerError.Msg("Internal server error"),
			errors.ErrInternalServerError.HTTPCode()))
	}
}

// getSpecialities godoc
// @Summary      Get Specialities
// @Description  Get all specialities of the system.
// @Tags         Specialities
// @Produce      json
// @Success      200  {object}  doctors.SpecialitiesResponse
// @Failure      500  {object}  we.ErrorResponse
// @Router       /doctors/v1/specialities [get]
func (h *DoctorsV1) getSpecialities(w http.ResponseWriter, r *http.Request) {

	especialities, err := h.app.GetSpecialities(r.Context())

	if err == nil {
		render.Render(w, r, NewSpecialitiesResponse(especialities))
	} else {
		render.Render(w, r, we.Err(
			errors.ErrInternalServerError.Msg("Internal server error"),
			errors.ErrInternalServerError.HTTPCode()))
	}
}
