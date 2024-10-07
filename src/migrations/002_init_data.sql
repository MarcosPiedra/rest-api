-- +goose Up
INSERT INTO doctors.specialty ("name") VALUES('Geriatric');
INSERT INTO doctors.specialty ("name") VALUES('Cardiologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Dermatologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Allergologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Anatomists');
INSERT INTO doctors.specialty ("name") VALUES('Andrologists');
INSERT INTO doctors.specialty ("name") VALUES('Anesthesiologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Gastroenterologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Geriatricians'); 
INSERT INTO doctors.specialty ("name") VALUES('Gynaecologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Immunologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Intensivists'); 
INSERT INTO doctors.specialty ("name") VALUES('Internists'); 
INSERT INTO doctors.specialty ("name") VALUES('Urologists');
INSERT INTO doctors.specialty ("name") VALUES('Virologists'); 
INSERT INTO doctors.specialty ("name") VALUES('Teratologists');
INSERT INTO doctors.specialty ("name") VALUES('Toxicologists');

INSERT INTO doctors.doctors
("name", surname, email, speciality_id, registration_id, phone, adress, city, zip_code, country)
VALUES('Marcos', 'Piedra Osuna', 'marcos@test.com', 1, '0001', '+34666123456', 'avenida diagonal', 'Barcelona', '08080', 'Spain');

INSERT INTO doctors.doctors
("name", surname, email, speciality_id, registration_id, phone, adress, city, zip_code, country)
VALUES('Juan', 'Pérez González', 'juan@test.com', 2, '0002', '+34636887744', 'calle mallorca', 'Barcelona', '08080', 'Spain');