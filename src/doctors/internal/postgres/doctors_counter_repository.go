package postgres

import (
	"context"
	"database/sql"
	"doctor-search-engine/doctors/internal/domain"

	"github.com/rs/zerolog"
	"github.com/stackus/errors"
)

type DoctorsCounterRepository struct {
	db     *sql.DB
	logger zerolog.Logger
}

// GetDoctorCounter implements domain.DoctorCounterRepository.
func (d *DoctorsCounterRepository) GetDoctorCounter(ctx context.Context) (doctors []*domain.DoctorCounter, err error) {

	query := `
	  SELECT 
	     d.id,
	     d."name",
         d.surname,
         d.speciality_id
      FROM doctors.doctor_counter dc
      INNER JOIN doctors.doctors d ON d.id = dc.doctor_id
	  ORDER BY dc.counter desc
	  LIMIT 3
	  `

	rows, err := d.db.QueryContext(ctx, query)

	if err != nil {
		d.logger.Err(err).Msg("")
		return nil, errors.Wrap(err, "querying counter")
	}

	defer func(rows *sql.Rows) {
		rows.Close()
	}(rows)

	for rows.Next() {
		doctor := &domain.DoctorCounter{}
		err := rows.Scan(&doctor.DoctorId,
			&doctor.Name,
			&doctor.Surname,
			&doctor.SpecialityId)
		if err != nil {
			return nil, errors.Wrap(err, "scanning counters row")
		}

		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing counters rows")
	}

	return doctors, nil
}

func NewDoctorsSearchCounter(db *sql.DB, logger zerolog.Logger) *DoctorsCounterRepository {
	return &DoctorsCounterRepository{
		db:     db,
		logger: logger,
	}
}

var _ domain.DoctorCounterRepository = (*DoctorsCounterRepository)(nil)
