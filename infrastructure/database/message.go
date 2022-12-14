package database

import (
	"context"
	"fmt"
	"time"

	"github.com/syougo1209/b-match-server/domain/model"
)

type MessageRepository struct {
	Db DbConnection
}

func (r *MessageRepository) FetchMessages(
	ctx context.Context, conversationID model.ConversationID,
	cursor, limit int,
) (model.Messages, error) {
	if limit < 0 {
		return nil, fmt.Errorf("limit is negative value")
	}

	dtos := []*Message{}
	sql := `SELECT
			id, send_user_id, conversation_id,
			type, text, created_at
		FROM message
		WHERE conversation_id =?
		AND id <= ?
		ORDER BY id desc
		LIMIT ?;`
	if err := r.Db.SelectContext(ctx, &dtos, sql, conversationID, cursor, limit); err != nil {
		return nil, fmt.Errorf("list conversation messages conversationID=%d: %v", conversationID, err)
	}
	messages := make(model.Messages, len(dtos))
	for i, v := range dtos {
		messages[i] = &model.Message{
			ID:             model.MessageID(v.ID),
			SendUserID:     model.UserID(v.SendUserID),
			ConversationID: model.ConversationID(v.ConversationID),
			Type:           model.MessageType(v.Type),
			Text:           v.Text,
			CreatedAt:      v.CreatedAt,
		}
	}
	return messages, nil
}

func (r *MessageRepository) CreateTextMessage(ctx context.Context, conversationID model.ConversationID, uid model.UserID, text string, now time.Time) (*model.Message, error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to start transaction")
	}

	query := `INSERT INTO message (send_user_id, conversation_id, type, text, created_at) VALUES (?,?,?,?,?);`
	result, err := tx.ExecContext(ctx, query, uid, conversationID, model.MessageTypeText, text, now)
	if err != nil {
		return nil, fmt.Errorf("failed to insert message: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("CreateTextMessage LastInsertID %w", err)
	}
	message := &model.Message{
		ID:             model.MessageID(id),
		SendUserID:     model.UserID(uid),
		ConversationID: model.ConversationID(conversationID),
		Type:           model.MessageTypeText,
		Text:           text,
		CreatedAt:      now,
	}
	return message, nil
}

type Message struct {
	ID             uint64    `db:"id"`
	SendUserID     uint64    `db:"send_user_id"`
	ConversationID uint64    `db:"conversation_id"`
	Type           uint      `db:"type"`
	Text           string    `db:"text"`
	CreatedAt      time.Time `db:"created_at"`
}
