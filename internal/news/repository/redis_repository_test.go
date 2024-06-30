package repository

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/internal/news"
	"github.com/AleksK1NG/api-mc/pkg/db/redis"
	"github.com/AleksK1NG/api-mc/pkg/logger"
)

func SetupRedis() (news.RedisRepository, func()) {
	path := "/tmp/db/" + uuid.NewString()
	client := redis.NewRedisClient(&config.Config{Redis: config.RedisConfig{
		RedisAddr: path,
	}, Session: config.Session{Expire: 3600}}, logger.NewApiLogger(&config.Config{Redis: config.RedisConfig{
		RedisAddr: "/tmp/db",
	}}))

	newsRedisRepo := NewNewsRedisRepo(client)

	if newsRedisRepo == nil {
		logger.NewApiLogger(&config.Config{Redis: config.RedisConfig{
			RedisAddr: "/tmp/db",
		}}).Errorf("Redis Client init: %s", errors.New(" error"))
	}

	return newsRedisRepo, func() {
		client.Close()
		os.RemoveAll(path)
	}
}

func TestNewsRedisRepo_SetNewsCtx(t *testing.T) {
	t.Parallel()

	newsRedisRepo, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("SetNewsCtx", func(t *testing.T) {
		newsUID := uuid.New()
		key := "key"
		n := &models.NewsBase{
			NewsID:  newsUID,
			Title:   "Title",
			Content: "Content",
		}

		err := newsRedisRepo.SetNewsCtx(context.Background(), key, 10, n)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

func TestNewsRedisRepo_GetNewsByIDCtx(t *testing.T) {
	t.Parallel()

	newsRedisRepo, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("GetNewsByIDCtx", func(t *testing.T) {
		newsUID := uuid.New()
		key := "key"
		n := &models.NewsBase{
			NewsID:  newsUID,
			Title:   "Title",
			Content: "Content",
		}

		newsBase, err := newsRedisRepo.GetNewsByIDCtx(context.Background(), key)
		require.Nil(t, newsBase)
		require.NotNil(t, err)

		err = newsRedisRepo.SetNewsCtx(context.Background(), key, 10, n)
		require.NoError(t, err)
		require.Nil(t, err)

		newsBase, err = newsRedisRepo.GetNewsByIDCtx(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
		require.NotNil(t, newsBase)
	})
}

func TestNewsRedisRepo_DeleteNewsCtx(t *testing.T) {
	t.Parallel()

	newsRedisRepo, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("SetNewsCtx", func(t *testing.T) {
		key := "key"

		err := newsRedisRepo.DeleteNewsCtx(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}
