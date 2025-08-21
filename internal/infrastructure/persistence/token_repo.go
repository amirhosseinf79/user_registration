package persistence

import (
	"context"
	"strconv"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/redis/go-redis/v9"
)

type tokenRepository struct {
	ctx             context.Context
	client          *redis.Client
	prefix          string
	refreshTokenExp time.Duration
}

func NewTokenRepository(ctx context.Context, client *redis.Client, refreshTokenExp time.Duration) repository.TokenRepository {
	return &tokenRepository{
		ctx:             ctx,
		client:          client,
		prefix:          "token:",
		refreshTokenExp: refreshTokenExp,
	}
}

func (t *tokenRepository) SaveRefreshToken(token *model.Token) error {
	return t.client.Set(t.ctx, t.prefix+token.RefreshToken, token.UserID, t.refreshTokenExp).Err()
}

func (t *tokenRepository) GetUserIDByRefresh(refreshToken string) (uint, error) {
	userID, err := t.client.Get(t.ctx, t.prefix+refreshToken).Result()
	if err == redis.Nil {
		return 0, dto.ErrObjectNotFound
	}
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
