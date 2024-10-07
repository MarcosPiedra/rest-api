package postgres

import (
	"context"
	"database/sql"
	"doctor-search-engine/doctors/internal/domain"

	"github.com/rs/zerolog"
	"github.com/stackus/errors"
)

type SpecialityRepository struct {
	db     *sql.DB
	logger zerolog.Logger
}

// GetSpeciality implements domain.SpecialityRepository.
func (s *SpecialityRepository) ExistsSpeciality(id int, ctx context.Context) (bool, error) {
	const query = `SELECT count(*) c from doctors.specialty d WHERE d.id = $1`
	row := s.db.QueryRowContext(ctx, query, id)

	if err := row.Err(); err != nil {
		return false, errors.Wrap(err, "Error quering")
	}

	count := 0
	row.Scan(&count)

	return count > 0, nil
}

func (s *SpecialityRepository) GetSpecialities(ctx context.Context) (specialities []*domain.Speciality, err error) {
	query := `
		SELECT  id,
			    name
		FROM doctors.specialty s`

	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		s.logger.Err(err).Msg("")
		return nil, errors.Wrap(err, "querying counter")
	}

	defer func(rows *sql.Rows) {
		rows.Close()
	}(rows)

	for rows.Next() {
		speciality := &domain.Speciality{}
		err := rows.Scan(&speciality.Id, &speciality.Name)
		if err != nil {
			return nil, errors.Wrap(err, "scanning counters row")
		}

		specialities = append(specialities, speciality)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing counters rows")
	}

	return specialities, nil
}

func NewSpecialityRepository(db *sql.DB, logger zerolog.Logger) *SpecialityRepository {
	return &SpecialityRepository{
		db:     db,
		logger: logger,
	}
}

var _ domain.SpecialityRepository = (*SpecialityRepository)(nil)
