package database

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestConversationRepository_UpdateLastMessageID(t *testing.T) {
	ctx := context.Background()

	db := testutils.NewDBForTest(t)
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx = context.WithValue(ctx, &txKey, tx)
	t.Cleanup(func() { _ = tx.Rollback() })

	messages, cid, _ := testutils.PrepareMessages(ctx, t, tx)

	repo := ConversationRepository{
		Db: tx,
	}
	tests := map[string]struct {
		cid     model.ConversationID
		mid     model.MessageID
		wantErr error
	}{
		"updateに成功すること": {cid, messages[0].ID, nil},
		"updateすべきデータが見つからない場合、エラーを返すこと": {model.ConversationID(100), messages[0].ID, model.ErrNotFound},
	}
	for n, tt := range tests {

		t.Run(n, func(t *testing.T) {
			err = repo.UpdateLastMessageID(ctx, tt.cid, tt.mid)
			if tt.wantErr == nil && err != nil {
				t.Fatalf("want no error, but got %v", err)
			}
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConversationRepository_FetchConversaionList(t *testing.T) {
	ctx := context.Background()

	db := testutils.NewDBForTest(t)
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = tx.Rollback() })

	repo := ConversationRepository{
		Db: tx,
	}

	tests := map[string]struct {
		uid           model.UserID
		conversations model.Conversations
	}{
		"ユーザーに紐付いた会話の一覧が取得できること": {},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			gotConversations, err = repo.FetchConversaionList(ctx, tt.uid)
			if err != nil {
				t.Fatalf("error during FetchConversaionList: %v", err)
			}
			if d := cmp.Diff(gotConversations, tt.conversations); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
