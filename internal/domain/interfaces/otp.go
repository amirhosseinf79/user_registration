package interfaces

import (
	auth_request "github.com/amirhosseinf79/user_registration/internal/dto/auth/request"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

type OTPStoreService interface {
	StoreCode(fields auth_request.FieldSendOTP) (string, *shared_dto.ResponseOneMessage)
	CheckOTPCode(fields auth_request.FieldVerifyOTP) (bool, *shared_dto.ResponseOneMessage)
}
