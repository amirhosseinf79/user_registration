package user

import "github.com/amirhosseinf79/user_registration/internal/dto/shared"

type ResponseList struct {
	Items []ResponseDetails     `json:"items"`
	Meta  shared.MetaPagination `json:"meta"`
}
