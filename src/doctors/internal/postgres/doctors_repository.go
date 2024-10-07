package postgres

import (
	"context"
	"database/sql"
	"doctor-search-engine/doctors/internal/domain"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/stackus/errors"
)

type DoctorsRepository struct {
	db     *sql.DB
	logger zerolog.Logger
}

func NewDoctorsRepository(db *sql.DB, logger zerolog.Logger) *DoctorsRepository {
	return &DoctorsRepository{
		db:     db,
		logger: logger,
	}
}

// ExistsRegistationId implements domain.DoctorRepository.
func (t *DoctorsRepository) ExistsRegistationId(ctx context.Context, registationId string) (bool, error) {

	const query = `select count(*) c from doctors.doctors d where d.registration_id = $1`
	row := t.db.QueryRowContext(ctx, query, registationId)

	if err := row.Err(); err != nil {
		return false, errors.Wrap(err, "Error quering")
	}

	count := 0
	row.Scan(&count)

	return count > 0, nil
}

func (t *DoctorsRepository) SearchDoctor(ctx context.Context, name string, surname string, specialityId int) (doctors []*domain.Doctor, err error) {

	var args = []any{}
	query := "WITH docs AS"
	query += "		(SELECT d.id,"
	query += `				d."name",`
	query += "				d.surname,"
	query += "				d.email,"
	query += "				d.speciality_id,"
	query += "				s.name speciality_name,"
	query += "				d.registration_id,"
	query += "				d.phone,"
	query += "				d.adress,"
	query += "				d.city,"
	query += "				d.zip_code,"
	query += "				d.country"
	query += "		FROM doctors.doctors d inner join doctors.specialty s on s.id = d.speciality_id"
	query += "		WHERE 1 = 1"

	count := 0
	if name != "" {
		count++
		args = append(args, "%"+name+"%")
		query += "		    AND d.name like $" + strconv.Itoa(count)
	}

	if surname != "" {
		count++
		args = append(args, "%"+surname+"%")
		query += "		    AND d.surname like $" + strconv.Itoa(count)
	}

	if specialityId != 0 {
		count++
		args = append(args, specialityId)
		query += "		    AND d.speciality_id = $" + strconv.Itoa(count)
	}

	query += "		),"
	query += "		insert_docs AS"
	query += "		(INSERT INTO doctors.doctor_counter (doctor_id, counter) SELECT id AS doctor_id,"
	query += "																		0 AS counter"
	query += "		FROM docs ON CONFLICT (doctor_id) DO UPDATE"
	query += "		SET counter = doctor_counter.counter + 1)"
	query += "		SELECT *"
	query += "		FROM docs"

	rows, err := t.db.QueryContext(ctx, query, args...)

	if err != nil {
		t.logger.Err(err).Msg("")
		return nil, errors.Wrap(err, "querying doctors")
	}

	defer func(rows *sql.Rows) {
		rows.Close()
	}(rows)

	for rows.Next() {
		doctor := &domain.Doctor{}
		err := rows.Scan(&doctor.Id,
			&doctor.Name,
			&doctor.Surname,
			&doctor.Email,
			&doctor.SpecialityId,
			&doctor.SpecialityName,
			&doctor.RegistrationId,
			&doctor.Phone,
			&doctor.Adress,
			&doctor.City,
			&doctor.ZipCode,
			&doctor.Country)
		if err != nil {
			return nil, errors.Wrap(err, "scanning doctors row")
		}

		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing doctors rows")
	}

	return doctors, nil
}

func (t *DoctorsRepository) AddDoctor(ctx context.Context, doctor *domain.Doctor) (int, error) {

	const query = `
		INSERT INTO doctors.doctors
					("name",
					surname,
					email,
					speciality_id,
					registration_id,
					phone,
					adress,
					city,
					zip_code,
					country)
		VALUES     ($1,
					$2,
					$3,
					$4,
					$5,
					$6,
					$7,
					$8,
					$9,
					$10)
		RETURNING id; 	
	`
	row := t.db.QueryRowContext(ctx, query,
		doctor.Name,
		doctor.Surname,
		doctor.Email,
		doctor.SpecialityId,
		doctor.RegistrationId,
		doctor.Phone,
		doctor.Adress,
		doctor.City,
		doctor.ZipCode,
		doctor.Country)

	if err := row.Err(); err != nil {
		t.logger.Err(err).Msg("")
		return 0, errors.Wrap(err, "creating doctors")
	}

	doctorId := 0
	row.Scan(&doctorId)

	return doctorId, nil
}

var _ domain.DoctorRepository = (*DoctorsRepository)(nil)
