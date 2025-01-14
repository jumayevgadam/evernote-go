package cache

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jumayevgadam/evernote-go/internal/connection"
	"github.com/jumayevgadam/evernote-go/internal/models/abstract"
)

var _ Store = (*ClientRedisRepo)(nil)

type Store interface {
	Get(ctx context.Context, argument abstract.CacheArgument) ([]byte, error)
	Set(ctx context.Context, argument abstract.CacheArgument, value []byte, duration time.Duration) error
	Del(ctx context.Context, argument abstract.CacheArgument) error
}

type ClientRedisRepo struct {
	rdb connection.Cache
}

func NewClientRDRepository(rdb connection.Cache) *ClientRedisRepo {
	return &ClientRedisRepo{rdb: rdb}
}

// getCacheKey from redisDB.
func (c *ClientRedisRepo) getCacheKey(objectType, id string) string {
	return strings.Join([]string{
		objectType,
		id,
	}, ":")
}

func (c *ClientRedisRepo) Get(ctx context.Context, argument abstract.CacheArgument) ([]byte, error) {
	key := argument.ToCacheStorage()
	cacheKey := c.getCacheKey(key.ObjectType, key.ID)

	valueString, err := c.rdb.Get(ctx, cacheKey)
	if err != nil {
		return nil, fmt.Errorf("error getting cachedValue: %w", err)
	}

	return []byte(valueString), nil
}

func (c *ClientRedisRepo) Set(ctx context.Context, argument abstract.CacheArgument, value []byte, duration time.Duration) error {
	key := argument.ToCacheStorage()
	cacheKey := c.getCacheKey(key.ObjectType, key.ID)

	err := c.rdb.Set(ctx, cacheKey, string(value), duration)
	if err != nil {
		return fmt.Errorf("error setting argument to cache: %w", err)
	}

	return nil
}

func (c *ClientRedisRepo) Del(ctx context.Context, argument abstract.CacheArgument) error {
	key := argument.ToCacheStorage()
	cacheKey := c.getCacheKey(key.ObjectType, key.ID)

	err := c.rdb.Del(ctx, cacheKey)
	if err != nil {
		return fmt.Errorf("error deleting cacheKey: %w", err)
	}

	return nil
}
