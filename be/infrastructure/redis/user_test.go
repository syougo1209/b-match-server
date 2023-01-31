package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestRedis_Save(t *testing.T) {
	cli := testutils.OpenRedisForTest(t)
	redis := &Redis{Cli: cli}
	key := "TestKVS_Save"
	uid := model.UserID(1)
	ctx := context.Background()
	t.Cleanup(func() {
		cli.Del(ctx, key)
	})
	if err := redis.Save(ctx, key, uid); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

func TestRedis_Load(t *testing.T) {
	cli := testutils.OpenRedisForTest(t)
	redis := &Redis{Cli: cli}
	ctx := context.Background()
	uid := model.UserID(1)
	key := "TestKVS_Load"

	tests := map[string]struct {
		prepareFn func()
		want      model.UserID
		wantErr   bool
	}{
		"loadするデータがあるとき": {
			func() {
				cli.Set(ctx, key, uint64(uid), 30*time.Minute)
				fmt.Print("halloe")
			},
			uid,
			false},
		"loadするデータがない時": {func() {}, model.UserID(0), true},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			tt.prepareFn()
			gots, err := redis.Load(ctx, key)
			if tt.wantErr == false && err != nil {
				t.Fatalf("want no error but got: %v", err)
			}
			if d := cmp.Diff(gots, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got, +want)\n%s", d)
			}
			t.Cleanup(func() {
				cli.Del(ctx, key)
			})
		})
	}
}
