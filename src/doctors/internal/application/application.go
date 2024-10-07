package application

import (
	"context"
	"doctor-search-engine/doctors/internal/application/commands"
	"doctor-search-engine/doctors/internal/application/queries"
	"doctor-search-engine/doctors/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		AddDoctor(ctx context.Context, cmd commands.AddDoctor) (*domain.Doctor, error)
	}

	Queries interface {
		SearchDoctors(ctx context.Context, query queries.SearchDoctor) ([]*domain.Doctor, error)
		GetDoctorsCounter(ctx context.Context) ([]*domain.DoctorCounter, error)
		GetSpecialities(ctx context.Context) ([]*domain.Speciality, error)
	}

	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.AddDoctorHandler
	}

	appQueries struct {
		queries.SearchDoctorHandler
		queries.GetDoctorCounterHandler
		queries.GetSpecialitiesHandler
	}
)

var _ App = (*Application)(nil)

func NewApplication(
	doctorRespository domain.DoctorRepository,
	specialityRepository domain.SpecialityRepository,
	doctorSearchCounterRepository domain.DoctorCounterRepository,
) *Application {
	return &Application{
		appCommands: appCommands{
			AddDoctorHandler: commands.NewAddDoctorHandler(doctorRespository, specialityRepository),
		},
		appQueries: appQueries{
			SearchDoctorHandler:     queries.NewSearchDoctorHandler(doctorRespository),
			GetSpecialitiesHandler:  queries.NewGetSpecialitiesHandler(specialityRepository),
			GetDoctorCounterHandler: queries.NewGetDoctorCounterHandler(doctorSearchCounterRepository),
		},
	}
}
