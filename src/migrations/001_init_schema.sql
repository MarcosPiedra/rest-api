-- +goose Up
CREATE SCHEMA doctors;

CREATE TABLE doctors.specialty
(
    id                 serial NOT NULL,
    name               text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE doctors.doctors
(
    id                 serial NOT NULL,
    name               text NOT NULL,
    surname            text NOT NULL,
    email              text NOT NULL,
    speciality_id      int NOT NULL REFERENCES doctors.specialty (id),
    registration_id    text NOT NULL,
    phone              text NOT NULL,
    adress             text NOT NULL,
    city               text NOT NULL,
    zip_code           text NOT NULL,
    country            text NOT null,
    PRIMARY KEY (id),
    UNIQUE(registration_id)
);

CREATE TABLE doctors.doctor_counter
(
    doctor_id          int NOT NULL REFERENCES doctors.doctors (id),
    counter            int NOT NULL,
    PRIMARY KEY (doctor_id)
);
