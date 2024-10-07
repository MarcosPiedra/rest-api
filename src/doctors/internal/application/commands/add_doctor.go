package commands

import (
	"context"
	"doctor-search-engine/doctors/internal/domain"
	"net/mail"

	err "github.com/stackus/errors"
)

type (
	AddDoctor struct {
		Name           string
		Surname        string
		Email          string
		RegistrationId string
		Phone          string
		Adress         string
		City           string
		ZipCode        string
		Country        string
		SpecialityId   int
	}

	AddDoctorHandler struct {
		doctorRepository     domain.DoctorRepository
		specialityRepository domain.SpecialityRepository
	}
)

func NewAddDoctorHandler(doctorRepository domain.DoctorRepository, specialityRepository domain.SpecialityRepository) AddDoctorHandler {
	return AddDoctorHandler{doctorRepository: doctorRepository, specialityRepository: specialityRepository}
}

func (h AddDoctorHandler) AddDoctor(ctx context.Context, cmd AddDoctor) (*domain.Doctor, error) {

	_, e := mail.ParseAddress(cmd.Email)

	if e != nil {
		return nil, err.Wrap(err.ErrBadRequest, "Invalid email")
	}

	existsExpecialityId, e := h.specialityRepository.ExistsSpeciality(cmd.SpecialityId, ctx)

	if e != nil {
		return nil, err.Wrap(err.ErrInternalServerError, e.Error())
	}

	if !existsExpecialityId {
		return nil, err.Wrap(err.ErrBadRequest, "Speciality id not registered")
	}

	existRegistationId, e := h.doctorRepository.ExistsRegistationId(ctx, cmd.RegistrationId)

	if e != nil {
		return nil, err.Wrap(err.ErrInternalServerError, e.Error())
	}

	if existRegistationId {
		return nil, err.Wrap(err.ErrConflict, "Doctor already exists with the same registation Id")
	}

	doctorToAdd := &domain.Doctor{
		Name:           cmd.Name,
		Surname:        cmd.Surname,
		Email:          cmd.Email,
		RegistrationId: cmd.RegistrationId,
		Phone:          cmd.Phone,
		Adress:         cmd.Adress,
		City:           cmd.City,
		ZipCode:        cmd.ZipCode,
		Country:        cmd.Country,
		SpecialityId:   cmd.SpecialityId,
	}

	doctorId, e := h.doctorRepository.AddDoctor(ctx, doctorToAdd)

	if e != nil {
		return nil, err.Wrap(err.ErrInternalServerError, e.Error())
	}

	if doctorId == 0 {
		return nil, err.Wrap(err.ErrInternalServerError, "Doctor not created")
	}

	doctorToAdd.Id = doctorId

	return doctorToAdd, nil
}
