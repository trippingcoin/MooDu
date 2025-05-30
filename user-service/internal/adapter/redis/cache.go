package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aftosmiros/moodu/user-service/internal/domain"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) GetUser(ctx context.Context, barcode string) (*domain.User, error) {
	key := fmt.Sprintf("user:%s", barcode)

	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // not in cache
	} else if err != nil {
		return nil, err
	}

	var user domain.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RedisCache) SetUser(ctx context.Context, user *domain.User, ttl time.Duration) error {
	key := fmt.Sprintf("user:%s", user.Barcode)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisCache) DeleteUser(ctx context.Context, barcode string) error {
	key := fmt.Sprintf("user:%s", barcode)
	return r.client.Del(ctx, key).Err()
}
