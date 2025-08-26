package user_response

import shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"

type List struct {
	Items []Details                 `json:"items"`
	Meta  shared_dto.MetaPagination `json:"meta"`
}
