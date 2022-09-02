package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/syougo1209/b-match-server/config"
	"github.com/syougo1209/b-match-server/domain/model"
	"golang.org/x/net/context"
)

func NewRedis(ctx context.Context, cfg *config.Config) (*Redis, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Redis{Cli: cli}, nil
}

type Redis struct {
	Cli *redis.Client
}

func (k *Redis) Save(ctx context.Context, key string, userID model.UserID) error {
	id := uint64(userID)
	return k.Cli.Set(ctx, key, id, 30*time.Minute).Err()
}

func (k *Redis) Load(ctx context.Context, key string) (model.UserID, error) {
	id, err := k.Cli.Get(ctx, key).Uint64()
	if err != nil {
		return 0, fmt.Errorf("failed to get by %q: %w", key, model.ErrNotFound)
	}
	return model.UserID(id), nil
}
