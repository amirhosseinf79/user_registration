package repository

import (
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type TokenRepository interface {
	Set(token *model.Token, exp time.Duration) error
	Get(token *model.Token) error
}
