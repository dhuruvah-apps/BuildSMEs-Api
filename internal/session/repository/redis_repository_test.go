package repository

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/internal/session"
	"github.com/AleksK1NG/api-mc/pkg/db/redis"
	"github.com/AleksK1NG/api-mc/pkg/logger"
)

func SetupRedis() (session.SessRepository, func()) {
	path := "/tmp/db/" + uuid.NewString()
	client := redis.NewRedisClient(&config.Config{Redis: config.RedisConfig{
		RedisAddr: path,
	}, Session: config.Session{Expire: 3600}}, logger.NewApiLogger(&config.Config{Redis: config.RedisConfig{
		RedisAddr: "/tmp/db",
	}}))

	sessRepository := NewSessionRepository(client, &config.Config{Redis: config.RedisConfig{
		RedisAddr: "/tmp/db",
	}})

	return sessRepository, func() {
		client.Close()
		os.RemoveAll(path)
	}
}

func TestSessionRepo_CreateSession(t *testing.T) {
	t.Parallel()

	sessRepository, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("CreateSession", func(t *testing.T) {
		sessUUID := uuid.New()
		sess := &models.Session{
			SessionID: sessUUID.String(),
			UserID:    sessUUID,
		}
		s, err := sessRepository.CreateSession(context.Background(), sess, 10)
		require.NoError(t, err)
		require.NotEqual(t, s, "")
	})
}

func TestSessionRepo_GetSessionByID(t *testing.T) {
	t.Parallel()

	sessRepository, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("GetSessionByID", func(t *testing.T) {
		sessUUID := uuid.New()
		sess := &models.Session{
			SessionID: sessUUID.String(),
			UserID:    sessUUID,
		}
		createdSess, err := sessRepository.CreateSession(context.Background(), sess, 10)
		require.NoError(t, err)
		require.NotEqual(t, createdSess, "")

		s, err := sessRepository.GetSessionByID(context.Background(), createdSess)
		require.NoError(t, err)
		require.NotEqual(t, s, "")
	})
}

func TestSessionRepo_DeleteByID(t *testing.T) {
	t.Parallel()

	sessRepository, clean := SetupRedis()
	defer t.Cleanup(func() { clean() })

	t.Run("DeleteByID", func(t *testing.T) {
		sessUUID := uuid.New()
		err := sessRepository.DeleteByID(context.Background(), sessUUID.String())
		require.NoError(t, err)
	})
}
