package phone_number_service

import (
	"go-jti/dto/payload"
	"go-jti/dto/response"
)

type PhoneNumberService interface {
	CreatePhoneNumber(payload payload.PhoneNumberPayload) (*response.PhoneNumberResponse, error)
	UpdatePhoneNumber(payload payload.UpdatePhoneNumberPayload, id string) (*response.PhoneNumberResponse, error)
	DeletePhoneNumber(id string) (*response.PhoneNumberResponse, error)
	FindPhoneNumberByID(id string) (*response.PhoneNumberResponse, error)
	FindOddPhoneNumber() (*[]response.PhoneNumberResponse, error)
	FindEvenPhoneNumber() (*[]response.PhoneNumberResponse, error)
	GenerateNumber() (*[]response.PhoneNumberResponse, error)
}
