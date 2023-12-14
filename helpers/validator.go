package helpers

// Validation code from: https://docs.gofiber.io/guide/validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

// This is the validator instance
// for more information see: https://github.com/go-playground/validator
var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (v XValidator) ValidateInput(c *fiber.Ctx, data interface{}) []error {
	if err := c.BodyParser(&data); err != nil {
		errorMessage := fmt.Sprintf("Could not parse JSON body: %s", err.Error())
		log.Error(errorMessage)
		return []error{c.Status(fiber.ErrBadRequest.Code).JSON(map[string]interface{}{"errors": []string{errorMessage}})}
	}
	if errs := v.Validate(data); len(errs) > 0 && errs[0].Error {
		errors := make([]string, 0)
		for _, err := range errs {
			errors = append(errors, fmt.Sprintf(
				"[%s]: '%v' | Failed validation '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		return []error{c.Status(fiber.ErrUnprocessableEntity.Code).JSON(map[string]interface{}{"errors": errors})}
	}
	return []error{}
}

var InputValidator = &XValidator{
	validator: validate,
}
