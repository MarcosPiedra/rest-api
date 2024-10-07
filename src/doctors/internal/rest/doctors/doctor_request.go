package doctors

type DoctorRequest struct {
	Name           string `json:"name" validate:"required,min=4,max=32"`
	Surname        string `json:"surname" validate:"required,min=4,max=32"`
	Email          string `json:"email" validate:"required,min=4,max=32"`
	RegistrationId string `json:"registration_id" validate:"required,min=4,max=32"`
	Phone          string `json:"phone" validate:"required,min=4,max=32"`
	Adress         string `json:"adress" validate:"required,min=4,max=32"`
	ZipCode        string `json:"zip_code" validate:"required,min=4,max=32"`
	Country        string `json:"country" validate:"required,min=4,max=32"`
	City           string `json:"city" validate:"required,min=4,max=32"`
	SpecialityId   int    `json:"speciality_id" validate:"required,gte=0,lte=99999"`
}
