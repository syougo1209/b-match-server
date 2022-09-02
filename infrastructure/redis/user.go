package redis

import (
	"fmt"
	"time"

	"github.com/syougo1209/b-match-server/domain/model"
	"golang.org/x/net/context"
)

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
