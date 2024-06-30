package repository

import (
	"context"
	"encoding/json"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/news"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/db/redis"
)

// News redis repository
type newsRedisRepo struct {
	redisClient *redis.BadgerStore
}

// News redis repository constructor
func NewNewsRedisRepo(redisClient *redis.BadgerStore) news.RedisRepository {
	return &newsRedisRepo{redisClient: redisClient}
}

// Get new by id
func (n *newsRedisRepo) GetNewsByIDCtx(ctx context.Context, key string) (*models.NewsBase, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "newsRedisRepo.GetNewsByIDCtx")
	defer span.Finish()

	newsBytes, err := n.redisClient.Get([]byte(key))
	if err != nil {
		return nil, errors.Wrap(err, "newsRedisRepo.GetNewsByIDCtx.redisClient.Get")
	}
	newsBase := &models.NewsBase{}
	if err = json.Unmarshal(newsBytes, newsBase); err != nil {
		return nil, errors.Wrap(err, "newsRedisRepo.GetNewsByIDCtx.json.Unmarshal")
	}

	return newsBase, nil
}

// Cache news item
func (n *newsRedisRepo) SetNewsCtx(ctx context.Context, key string, seconds int, news *models.NewsBase) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "newsRedisRepo.SetNewsCtx")
	defer span.Finish()

	newsBytes, err := json.Marshal(news)
	if err != nil {
		return errors.Wrap(err, "newsRedisRepo.SetNewsCtx.json.Marshal")
	}
	if err = n.redisClient.Set([]byte(key), newsBytes); err != nil {
		return errors.Wrap(err, "newsRedisRepo.SetNewsCtx.redisClient.Set")
	}
	return nil
}

// Delete new item from cache
func (n *newsRedisRepo) DeleteNewsCtx(ctx context.Context, key string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "newsRedisRepo.DeleteNewsCtx")
	defer span.Finish()

	if err := n.redisClient.Del([]byte(key)); err != nil {
		return errors.Wrap(err, "newsRedisRepo.DeleteNewsCtx.redisClient.Del")
	}
	return nil
}
