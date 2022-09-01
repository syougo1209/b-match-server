package database

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestMessageRepository_FetchMessages(t *testing.T) {
	ctx := context.Background()

	db := testutils.NewDBForTest(t)
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = tx.Rollback() })
	repo := MessageRepository{
		Db: tx,
	}
	messages, cid, maxID := testutils.PrepareMessages(ctx, t, tx)
	wants := model.Messages{messages[1], messages[0]}

	type args struct {
		cursor, limit int
	}
	tests := map[string]struct {
		args args
		want model.Messages
	}{
		"max_id以下のidのメッセージが存在する場合、配列が取得できること": {args{maxID, 15}, wants},
		"max_idが0の時、空の配列が得られること":              {args{0, 15}, model.Messages{}},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			gots, err := repo.FetchMessages(ctx, cid, tt.args.cursor, tt.args.limit)
			if err != nil {
				t.Fatalf("error during ListConversationMessages: %v", err)
			}
			if d := cmp.Diff(gots, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got, +want)\n%s", d)
			}
		})
	}
}
