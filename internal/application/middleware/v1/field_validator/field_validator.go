package field_validator

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type fieldsValidatorMiddleware struct{}

func NewFieldValidator() interfaces.FieldValidatorMiddleware {
	return &fieldsValidatorMiddleware{}
}
