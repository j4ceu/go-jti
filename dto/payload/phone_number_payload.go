package payload

type PhoneNumberPayload struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Provider    string `json:"provider" validate:"required"`
}

type UpdatePhoneNumberPayload struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}
