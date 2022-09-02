package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/syougo1209/b-match-server/config"
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
