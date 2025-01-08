package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/jumayevgadam/evernote-go/internal/config"
	"github.com/redis/go-redis/v9"
)

// Ensure Redis struct implements the Cache interface.
var _ Cache = (*Redis)(nil)

// Redis struct keeps redis connection.
type Redis struct {
	redis *redis.Client
}

// NewCache returns a new Redis struct.
func NewCache(ctx context.Context, cfgs config.RedisDB) (*Redis, error) {
	options := &redis.Options{
		Addr:     cfgs.Address,
		Password: cfgs.Password,
		DB:       0,
	}

	rdb := redis.NewClient(options)
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping: %w", err)
	}

	return &Redis{redis: rdb}, nil
}

// Cache interface with Get, Set, Del and Close methods.
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	Close() error
}

// Get method fetches needed value using key.
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

// Set method insert a new key value pair in redisDB.
func (r *Redis) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.redis.Set(ctx, key, value, expiration).Err()
}

// Del method deletes the key value pair using identified key.
func (r *Redis) Del(ctx context.Context, key string) error {
	return r.redis.Del(ctx, key).Err()
}

// Close closes redisDB.
func (r *Redis) Close() error {
	return r.redis.Close()
}
