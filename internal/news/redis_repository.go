//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package news

import (
	"context"

	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
)

// News redis repository
type RedisRepository interface {
	GetNewsByIDCtx(ctx context.Context, key string) (*models.NewsBase, error)
	SetNewsCtx(ctx context.Context, key string, seconds int, news *models.NewsBase) error
	DeleteNewsCtx(ctx context.Context, key string) error
}
