package phone_number_service

import (
	"errors"
	"go-jti/dto/payload"
	"go-jti/dto/response"
	"go-jti/models"
	"go-jti/repositories/phone_number_repository"
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
	phoneNumber := models.PhoneNumber{
		PhoneNumber: payload.PhoneNumber,
		Provider:    payload.Provider,
	}

	phoneNumber, err := s.repository.CreatePhoneNumber(phoneNumber)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewPhoneNumberResponse(phoneNumber), nil
}

func (s *phoneNumberService) UpdatePhoneNumber(payload payload.PhoneNumberPayload, id string) (*response.PhoneNumberResponse, error) {
	phoneNumber, err := s.repository.FindPhoneNumberByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	if payload.PhoneNumber != "" {
		phoneNumber.PhoneNumber = payload.PhoneNumber
	}

	if payload.Provider != "" {
		phoneNumber.Provider = payload.Provider
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
		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("204")
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
