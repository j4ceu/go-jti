package phone_number_repository

import "go-jti/models"

type PhoneNumberRepository interface {
	CreatePhoneNumber(phoneNumber models.PhoneNumber) (models.PhoneNumber, error)
	UpdatePhoneNumber(phoneNumber models.PhoneNumber, id string) (models.PhoneNumber, error)
	DeletePhoneNumber(id string) (models.PhoneNumber, error)
	FindPhoneNumberByID(id string) (models.PhoneNumber, error)
	FindOddPhoneNumber() ([]models.PhoneNumber, error)
	FindEvenPhoneNumber() ([]models.PhoneNumber, error)
}
