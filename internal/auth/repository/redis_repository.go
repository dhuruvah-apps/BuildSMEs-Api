package repository

import (
	"context"
	"encoding/json"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/auth"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/db/redis"
)

// Auth redis repository
type authRedisRepo struct {
	redisClient *redis.BadgerStore
}

// Auth redis repository constructor
func NewAuthRedisRepo(redisClient *redis.BadgerStore) auth.RedisRepository {
	return &authRedisRepo{redisClient: redisClient}
}

// Get user by id
func (a *authRedisRepo) GetByIDCtx(ctx context.Context, key string) (*models.User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "authRedisRepo.GetByIDCtx")
	defer span.Finish()

	userBytes, err := a.redisClient.Get([]byte(key))
	if err != nil {
		return nil, errors.Wrap(err, "authRedisRepo.GetByIDCtx.redisClient.Get")
	}
	user := &models.User{}
	if err = json.Unmarshal(userBytes, user); err != nil {
		return nil, errors.Wrap(err, "authRedisRepo.GetByIDCtx.json.Unmarshal")
	}
	return user, nil
}

// Cache user with duration in seconds
func (a *authRedisRepo) SetUserCtx(ctx context.Context, key string, seconds int, user *models.User) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "authRedisRepo.SetUserCtx")
	defer span.Finish()

	userBytes, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "authRedisRepo.SetUserCtx.json.Unmarshal")
	}
	if err = a.redisClient.Set([]byte(key), userBytes); err != nil {
		return errors.Wrap(err, "authRedisRepo.SetUserCtx.redisClient.Set")
	}
	return nil
}

// Delete user by key
func (a *authRedisRepo) DeleteUserCtx(ctx context.Context, key string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "authRedisRepo.DeleteUserCtx")
	defer span.Finish()

	if err := a.redisClient.Del([]byte(key)); err != nil {
		return errors.Wrap(err, "authRedisRepo.DeleteUserCtx.redisClient.Del")
	}
	return nil
}
