package doctors

import (
	"doctor-search-engine/doctors/internal/application"
	"doctor-search-engine/doctors/internal/postgres"
	"doctor-search-engine/doctors/internal/rest"
	"doctor-search-engine/doctors/internal/rest/doctors"
	"doctor-search-engine/internal/system"
)

type Module struct {
	system *system.System
	api    *rest.Api
	app    application.App
}

func NewModule(system *system.System, api *rest.Api, app application.App) *Module {
	return &Module{
		system: system,
		api:    api,
		app:    app,
	}
}

func Build(s *system.System) *Module {
	mux := s.Mux()
	db := s.Db()
	logger := s.Logger()

	doctorsRepository := postgres.NewDoctorsRepository(db, logger)
	doctorsSearchCounterRepository := postgres.NewDoctorsSearchCounter(db, logger)
	specialityRepository := postgres.NewSpecialityRepository(db, logger)

	application := application.NewApplication(doctorsRepository, specialityRepository, doctorsSearchCounterRepository)
	doctorsV1 := doctors.NewDoctorsV1(application)
	api := rest.NewApi(mux, doctorsV1)
	module := NewModule(s, api, application)
	return module
}

// Start implements system.Module.
func (m *Module) Start() {
	m.api.Init()
}

var _ system.Module = (*Module)(nil)
