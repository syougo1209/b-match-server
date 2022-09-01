package database

import (
	"context"
	"testing"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestConversationStateRepository_ReadMessages(t *testing.T) {
	ctx := context.Background()

	db := testutils.NewDBForTest(t)
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = tx.Rollback() })

	user := testutils.PrepareUser(ctx, t, tx)
	otherUser := testutils.PrepareUser(ctx, t, tx)
	cid := testutils.PrepareConversation(ctx, t, tx)
	_ = testutils.PrepareConversationState(ctx, t, tx, cid, *user, *otherUser)

	type args struct {
		conversationID model.ConversationID
		messageID      model.MessageID
	}
	tests := map[string]struct {
		args    args
		wantErr bool
	}{
		"conversationが存在するとき、更新に成功すること":  {args{cid, model.MessageID(1)}, false},
		"conversationが存在しないとき、エラーが変えること": {args{model.ConversationID(100), model.MessageID(1)}, true},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			repo := ConversationStateRepository{
				Db: tx,
			}
			err := repo.ReadMessages(ctx, user.ID, tt.args.conversationID, tt.args.messageID)
			if tt.wantErr == false && err != nil {
				t.Fatalf("error during ListConversationMessages: %v", err)
			}
		})
	}
}
