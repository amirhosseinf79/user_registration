package user_request

import shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"

type FilterUser struct {
	FieldEmail
	PhoneNumber string `query:"phoneNumber"`
	Name        string `query:"name"`
	shared_dto.FieldPagination
}
