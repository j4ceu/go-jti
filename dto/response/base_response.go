package response

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors"`
	Status  int         `json:"status"`
}

type EmptyObj struct{}

func ConvertToBaseResponse(message string, status int, data interface{}) BaseResponse {
	return BaseResponse{
		Message: message,
		Data:    data,
		Status:  status,
		Errors:  nil,
	}
}

func ConvertErrorToBaseResponse(message string, status int, data interface{}, err string) BaseResponse {
	splittedError := strings.Split(err, "\n")
	return BaseResponse{
		Message: message,
		Data:    data,
		Status:  status,
		Errors:  splittedError,
	}
}

func Success(ctx *fiber.Ctx, message string, status int, data interface{}) error {
	response := ConvertToBaseResponse(message, status, data)
	return ctx.Status(status).JSON(response)
}

func SuccessWithPagination(ctx *fiber.Ctx, message string, status int, data interface{}) error {
	response := BaseResponse{
		Message: message,
		Data:    data,
		Status:  status,
		Errors:  nil,
	}
	return ctx.Status(status).JSON(response)
}

func Error(ctx *fiber.Ctx, message string, status int, err error) error {
	response := ConvertErrorToBaseResponse(message, status, EmptyObj{}, err.Error())
	return ctx.Status(status).JSON(response)
}

func ErrorValidate(ctx *fiber.Ctx, message string, status int, err interface{}) error {
	response := BaseResponse{
		Message: message,
		Data:    EmptyObj{},
		Status:  status,
		Errors:  err,
	}
	return ctx.Status(status).JSON(response)
}
