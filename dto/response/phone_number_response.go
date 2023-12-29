package response

import "go-jti/models"

type PhoneNumberResponse struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Provider    string `json:"provider"`
}

func NewPhoneNumberResponse(phoneNumber models.PhoneNumber) *PhoneNumberResponse {
	response := PhoneNumberResponse{
		ID:          phoneNumber.ID.String(),
		PhoneNumber: phoneNumber.PhoneNumber,
		Provider:    phoneNumber.Provider,
	}

	return &response
}

func NewPhoneNumberResponses(phoneNumbers []models.PhoneNumber) *[]PhoneNumberResponse {
	var responses []PhoneNumberResponse

	for _, phoneNumber := range phoneNumbers {
		responses = append(responses, *NewPhoneNumberResponse(phoneNumber))
	}

	return &responses
}
