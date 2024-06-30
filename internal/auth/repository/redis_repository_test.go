package repository

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/auth"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/db/redis"
	"github.com/AleksK1NG/api-mc/pkg/logger"
)

func SetupRedis() (auth.RedisRepository, func()) {
	path := "/tmp/db/" + uuid.NewString()
	client := redis.NewRedisClient(&config.Config{Redis: config.RedisConfig{
		RedisAddr: path,
	}, Session: config.Session{Expire: 3600}}, logger.NewApiLogger(&config.Config{Redis: config.RedisConfig{
		RedisAddr: "/tmp/db",
	}}))

	authRedisRepo := NewAuthRedisRepo(client)

	return authRedisRepo, func() {
		client.Close()
		os.RemoveAll(path)
	}
}

func TestAuthRedisRepo_GetByIDCtx(t *testing.T) {
	t.Parallel()

	authRedisRepo, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("GetByIDCtx", func(t *testing.T) {
		key := uuid.New().String()
		userID := uuid.New()
		u := &models.User{
			UserID:    userID,
			FirstName: "Alex",
			LastName:  "Bryksin",
		}

		err := authRedisRepo.SetUserCtx(context.Background(), key, 10, u)
		require.NoError(t, err)
		require.Nil(t, err)

		user, err := authRedisRepo.GetByIDCtx(context.Background(), key)
		require.NoError(t, err)
		require.NotNil(t, user)
	})
}

func TestAuthRedisRepo_SetUserCtx(t *testing.T) {
	t.Parallel()

	authRedisRepo, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("SetUserCtx", func(t *testing.T) {

		key := uuid.New().String()
		userID := uuid.New()
		u := &models.User{
			UserID:    userID,
			FirstName: "Alex",
			LastName:  "Bryksin",
		}

		err := authRedisRepo.SetUserCtx(context.Background(), key, 10, u)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

func TestAuthRedisRepo_DeleteUserCtx(t *testing.T) {
	t.Parallel()

	authRedisRepo, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("DeleteUserCtx", func(t *testing.T) {
		key := uuid.New().String()

		err := authRedisRepo.DeleteUserCtx(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}
