package utils

import (
	"go-jti/dto/response"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(payload any) []*response.ErrorResponse {
	validate := validator.New()
	var errors []*response.ErrorResponse

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
