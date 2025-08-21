package persistence

import (
	"context"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type tokenRepository struct {
	ctx             context.Context
	client          *redis.Client
	prefix          string
	refreshRokenExp time.Duration
}

func NewTokenRepository(ctx context.Context, client *redis.Client, refreshRokenExp time.Duration) repository.TokenRepository {
	return &tokenRepository{
		ctx:             ctx,
		client:          client,
		prefix:          "token:",
		refreshRokenExp: refreshRokenExp,
	}
}

func (t *tokenRepository) Set(token *model.Token) error {
	return t.client.Set(t.ctx, t.prefix+token.RefreshToken, token.AccessToken, t.refreshRokenExp).Err()
}

func (t *tokenRepository) Get(token *model.Token) error {
	accessToken, err := t.client.Get(t.ctx, t.prefix+token.RefreshToken).Result()
	if err != redis.Nil {
		return err
	}
	token.AccessToken = accessToken
	return nil
}
