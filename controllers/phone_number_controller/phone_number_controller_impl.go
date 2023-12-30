package phone_number_controller

import (
	"errors"
	"go-jti/dto/payload"
	"go-jti/dto/response"
	"go-jti/services/phone_number_service"
	"go-jti/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type phoneNumberController struct {
	service phone_number_service.PhoneNumberService
}

func NewPhoneNumberController(service phone_number_service.PhoneNumberService) PhoneNumberController {
	return &phoneNumberController{service: service}
}

func (c *phoneNumberController) CreatePhoneNumber(ctx *fiber.Ctx) error {
	var payload payload.PhoneNumberPayload

	if err := ctx.BodyParser(&payload); err != nil {
		return response.Error(ctx, "failed", http.StatusBadRequest, err)
	}

	if err := utils.ValidateStruct(payload); err != nil {
		return response.ErrorValidate(ctx, "failed", http.StatusBadRequest, err)
	}

	phoneNumber, err := c.service.CreatePhoneNumber(payload)
	if err != nil {
		return response.Error(ctx, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(ctx, "success", http.StatusOK, phoneNumber)

}

func (c *phoneNumberController) UpdatePhoneNumber(ctx *fiber.Ctx) error {
	var payload payload.UpdatePhoneNumberPayload

	if err := ctx.BodyParser(&payload); err != nil {
		return response.Error(ctx, "failed", http.StatusBadRequest, err)
	}

	if err := utils.ValidateStruct(payload); err != nil {
		return response.ErrorValidate(ctx, "failed", http.StatusBadRequest, err)
	}

	id := ctx.Params("id")

	phoneNumber, err := c.service.UpdatePhoneNumber(payload, id)
	if err != nil {
		if err.Error() == "204" {
			return response.Error(ctx, "failed", http.StatusNoContent, errors.New("not found"))
		}
		return response.Error(ctx, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(ctx, "success", http.StatusOK, phoneNumber)
}

func (c *phoneNumberController) DeletePhoneNumber(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	phoneNumber, err := c.service.DeletePhoneNumber(id)
	if err != nil {
		return response.Error(ctx, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(ctx, "success", http.StatusOK, phoneNumber)

}

func (c *phoneNumberController) FindPhoneNumberByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	phoneNumber, err := c.service.FindPhoneNumberByID(id)
	if err != nil {
		return response.Error(ctx, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(ctx, "success", http.StatusOK, phoneNumber)
}

func (c *phoneNumberController) FindOddPhoneNumber(ctx *fiber.Ctx) error {
	phoneNumber, err := c.service.FindOddPhoneNumber()
	if err != nil {
		return response.Error(ctx, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(ctx, "success", http.StatusOK, phoneNumber)
}

func (c *phoneNumberController) FindEvenPhoneNumber(ctx *fiber.Ctx) error {
	phoneNumber, err := c.service.FindEvenPhoneNumber()
	if err != nil {
		return response.Error(ctx, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(ctx, "success", http.StatusOK, phoneNumber)
}

func (c *phoneNumberController) GeneratePhoneNumber(ctx *fiber.Ctx) error {
	phoneNumber, err := c.service.GenerateNumber()
	if err != nil {
		return response.Error(ctx, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(ctx, "success", http.StatusOK, phoneNumber)
}
