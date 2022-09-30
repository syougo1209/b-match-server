package testutils

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/b-match-server/domain/model"
)

func PrepareUser(ctx context.Context, t *testing.T, db *sqlx.Tx) *model.User {
	t.Helper()
	u := &model.User{
		Name:      "example",
		CreatedAt: time.Now(),
	}
	result, err := db.ExecContext(ctx, "INSERT INTO user (name, created_at) VALUES (?, ?)",
		u.Name, u.CreatedAt)
	if err != nil {
		t.Fatalf("error insert user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("error got user_id: %v", err)
	}

	u.ID = model.UserID(id)
	return u
}

func PrepareConversation(ctx context.Context, t *testing.T, db *sqlx.Tx) model.ConversationID {
	t.Helper()

	result, err := db.ExecContext(ctx, "INSERT INTO conversation (last_message_id) VALUES (?)", 0)
	if err != nil {
		t.Fatalf("error insert conversation: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("error got conversationID: %v", err)
	}

	return model.ConversationID(id)
}

func PrepareConversationState(
	ctx context.Context, t *testing.T, db *sqlx.Tx,
	cid model.ConversationID, fromUser model.User, toUser model.User,
) *model.Conversation {
	t.Helper()
	c := model.Conversation{
		ID:                  cid,
		FromUser:            &fromUser,
		ToUser:              toUser,
		UnreadMessagesCount: 0,
		LastMessage:         nil,
	}
	_, err := db.ExecContext(
		ctx,
		"INSERT INTO conversation_state (conversation_id, from_user_id, to_user_id, unread_messages_count, last_read_message_id) VALUES (?, ?, ?, ?, ?)",
		c.ID, c.FromUser.ID, c.ToUser.ID,
		c.UnreadMessagesCount, 0,
	)
	if err != nil {
		t.Fatalf("error insert conversation: %v", err)
	}
	return &c
}

func PrepareMessages(ctx context.Context, t *testing.T, db *sqlx.Tx) (model.Messages, model.ConversationID, int) {
	t.Helper()
	user := PrepareUser(ctx, t, db)
	conversationID := PrepareConversation(ctx, t, db)
	otherConversationID := PrepareConversation(ctx, t, db)

	wants := model.Messages{
		{
			SendUserID:     user.ID,
			ConversationID: conversationID,
			Type:           1,
			Text:           "want task 1",
			CreatedAt:      time.Now(),
		},
		{
			SendUserID:     user.ID,
			ConversationID: conversationID,
			Type:           1,
			Text:           "want task 2",
			CreatedAt:      time.Now(),
		},
	}

	messages := model.Messages{
		wants[0],
		{
			SendUserID:     user.ID,
			ConversationID: otherConversationID,
			Type:           1,
			Text:           "not want task",
			CreatedAt:      time.Now(),
		},
		wants[1],
	}

	result, err := db.ExecContext(ctx,
		`INSERT INTO message (send_user_id, conversation_id, type, text, created_at)
			VALUES
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?);`,
		messages[0].SendUserID, messages[0].ConversationID, messages[0].Type, messages[0].Text, messages[0].CreatedAt,
		messages[1].SendUserID, messages[1].ConversationID, messages[1].Type, messages[1].Text, messages[1].CreatedAt,
		messages[2].SendUserID, messages[2].ConversationID, messages[2].Type, messages[2].Text, messages[2].CreatedAt,
	)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		t.Fatal(err)
	}
	messages[0].ID = model.MessageID(id)
	messages[1].ID = model.MessageID(id + 1)
	messages[2].ID = model.MessageID(id + 2)

	return wants, model.ConversationID(conversationID), int(id + 2)
}
