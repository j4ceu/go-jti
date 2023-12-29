package phone_number_repository

import (
	"go-jti/models"

	"gorm.io/gorm"
)

type phoneNumberRepository struct {
	db *gorm.DB
}

func NewPhoneNumberRepository(db *gorm.DB) PhoneNumberRepository {
	return &phoneNumberRepository{db: db}
}

func (r *phoneNumberRepository) CreatePhoneNumber(phoneNumber models.PhoneNumber) (models.PhoneNumber, error) {
	err := r.db.Create(&phoneNumber).Error

	return phoneNumber, err
}

func (r *phoneNumberRepository) UpdatePhoneNumber(phoneNumber models.PhoneNumber, id string) (models.PhoneNumber, error) {
	err := r.db.Where("id = ?", id).Updates(&phoneNumber).Error

	if err != nil {
		return models.PhoneNumber{}, err
	}

	return r.FindPhoneNumberByID(id)

}

func (r *phoneNumberRepository) DeletePhoneNumber(id string) (models.PhoneNumber, error) {
	err := r.db.Where("id = ?", id).Delete(models.PhoneNumber{}).Error

	return models.PhoneNumber{}, err
}

func (r *phoneNumberRepository) FindPhoneNumberByID(id string) (models.PhoneNumber, error) {
	var phoneNumber models.PhoneNumber
	err := r.db.First(&phoneNumber, "id = ?", id).Error

	return phoneNumber, err
}

func (r *phoneNumberRepository) FindOddPhoneNumber() ([]models.PhoneNumber, error) {
	var phoneNumber []models.PhoneNumber
	err := r.db.Where("mod(cast(phone_number as bigint),2) <> 0").Find(&phoneNumber).Error

	return phoneNumber, err
}

func (r *phoneNumberRepository) FindEvenPhoneNumber() ([]models.PhoneNumber, error) {
	var phoneNumber []models.PhoneNumber
	err := r.db.Where("mod(cast(phone_number as bigint),2) = 0").Find(&phoneNumber).Error

	return phoneNumber, err
}
