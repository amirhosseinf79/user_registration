package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type tokenRepository struct {
	ctx    context.Context
	client *redis.Client
	prefix string
}

func NewTokenRepository(ctx context.Context, client *redis.Client) repository.TokenRepository {
	return &tokenRepository{
		ctx:    ctx,
		client: client,
		prefix: "token:",
	}
}

func (t *tokenRepository) Set(token *model.Token, exp time.Duration) error {
	return t.client.Set(t.ctx, t.prefix+token.UserID, token, exp).Err()
}

func (t *tokenRepository) Get(token *model.Token) error {
	errStr := t.client.HGetAll(t.ctx, t.prefix+token.UserID).Scan(token).Error()
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}
