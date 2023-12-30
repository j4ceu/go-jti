package phone_number_service

import (
	"errors"
	"go-jti/dto/payload"
	"go-jti/dto/response"
	"go-jti/models"
	"go-jti/repositories/phone_number_repository"
	"go-jti/utils"
	"log"

	"gorm.io/gorm"
)

type phoneNumberService struct {
	repository phone_number_repository.PhoneNumberRepository
}

func NewPhoneNumberService(repository phone_number_repository.PhoneNumberRepository) PhoneNumberService {
	return &phoneNumberService{repository: repository}
}

func (s *phoneNumberService) CreatePhoneNumber(payload payload.PhoneNumberPayload) (*response.PhoneNumberResponse, error) {

	decryptPhoneNumber, err := utils.GetAESDecrypted(payload.PhoneNumber)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	decryptProvider, err := utils.GetAESDecrypted(payload.Provider)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	phoneNumber := models.PhoneNumber{
		PhoneNumber: string(decryptPhoneNumber),
		Provider:    string(decryptProvider),
	}

	phoneNumber, err = s.repository.CreatePhoneNumber(phoneNumber)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewPhoneNumberResponse(phoneNumber), nil
}

func (s *phoneNumberService) UpdatePhoneNumber(payload payload.UpdatePhoneNumberPayload, id string) (*response.PhoneNumberResponse, error) {
	phoneNumber, err := s.repository.FindPhoneNumberByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	if payload.PhoneNumber != "" {
		decryptPhoneNumber, err := utils.GetAESDecrypted(payload.PhoneNumber)
		if err != nil {
			log.Println(string("\033[31m"), err.Error())
			return nil, err
		}
		phoneNumber.PhoneNumber = string(decryptPhoneNumber)
	}

	phoneNumber, err = s.repository.UpdatePhoneNumber(phoneNumber, id)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewPhoneNumberResponse(phoneNumber), nil
}

func (s *phoneNumberService) DeletePhoneNumber(id string) (*response.PhoneNumberResponse, error) {
	phoneNumber, err := s.repository.DeletePhoneNumber(id)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewPhoneNumberResponse(phoneNumber), nil
}

func (s *phoneNumberService) FindPhoneNumberByID(id string) (*response.PhoneNumberResponse, error) {
	phoneNumber, err := s.repository.FindPhoneNumberByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	return response.NewPhoneNumberResponse(phoneNumber), nil
}

func (s *phoneNumberService) FindOddPhoneNumber() (*[]response.PhoneNumberResponse, error) {
	phoneNumber, err := s.repository.FindOddPhoneNumber()
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("204")
	}

	return response.NewPhoneNumberResponses(phoneNumber), nil
}

func (s *phoneNumberService) FindEvenPhoneNumber() (*[]response.PhoneNumberResponse, error) {
	phoneNumber, err := s.repository.FindEvenPhoneNumber()
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("204")
	}

	return response.NewPhoneNumberResponses(phoneNumber), nil
}

func (s *phoneNumberService) GenerateNumber() (*[]response.PhoneNumberResponse, error) {
	var phoneNumbers []models.PhoneNumber

	for i := 0; i < 25; i++ {
		check := true
		number, provider := utils.GenerateRandomNumber()
		for check {
			exists, err := s.repository.CheckPhoneNumberExists(number)
			if err != nil {
				return nil, err
			}
			if !exists {
				phoneNumber := models.PhoneNumber{
					PhoneNumber: number,
					Provider:    provider,
				}
				phoneNumbers = append(phoneNumbers, phoneNumber)
				check = false
			}
		}
	}
	return response.NewPhoneNumberResponses(phoneNumbers), nil

}
