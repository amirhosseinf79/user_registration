package repository

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type TokenRepository interface {
	Set(token *model.Token) error
	Get(token *model.Token) error
}
