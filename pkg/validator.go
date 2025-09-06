package pkg

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func validateStruct[T any](request *T) (map[string]string, error) {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		errorsMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorsMap[err.Field()] = fmt.Sprintf("Field '%s' failed validation with tag '%s'", err.Field(), err.Tag())
		}
		return errorsMap, err
	}
	return nil, nil
}

func ValidateRequestBody[T any](request *T, c *fiber.Ctx) (map[string]string, error) {
	if err := c.BodyParser(request); err != nil {
		return map[string]string{
			"error": "Invalid request data",
		}, errors.New("invalid request data")
	}
	return validateStruct(request)
}

func ValidateQueryParams[T any](request *T, c *fiber.Ctx) (map[string]string, error) {
	if err := c.QueryParser(request); err != nil {
		return map[string]string{
			"error": "Invalid request data",
		}, errors.New("invalid request data")
	}
	return validateStruct(request)
}
