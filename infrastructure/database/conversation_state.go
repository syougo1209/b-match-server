package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
)

type ConversationStateRepository struct {
	Db DbConnection
}

func (csr *ConversationStateRepository) ReadMessages(
	ctx context.Context, uid model.UserID,
	conversationID model.ConversationID, messageID model.MessageID,
) error {
	var dto ConversationState
	query := `
	  select *
	  from conversation_state
		WHERE conversation_id =? AND from_user_id =?
	`
	if err := csr.Db.GetContext(ctx, &dto, query, conversationID, uid); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("GetContext conversation_state by conversation_id=%d, from_user_id=%d: %w", conversationID, uid, model.ErrNotFound)
		}
		return fmt.Errorf("GetContext conversation_state by conversation_id=%d, from_user_id=%d: %w", conversationID, uid, err)
	}

	uQuery := `
		UPDATE
		conversation_state
		SET unread_messages_count=?, last_read_message_id=?
		WHERE conversation_id =? AND from_user_id =?
	`

	_, err := csr.Db.ExecContext(ctx, uQuery, 0, messageID, conversationID, uid)
	if err != nil {
		return fmt.Errorf("update conversation_state conversation_id=%d from_user_id=%d: %w", conversationID, uid, err)
	}
	return nil
}

type ConversationState struct {
	ConversationID      uint64 `db:"conversation_id"`
	FromUserID          uint64 `db:"from_user_id"`
	ToUserID            uint64 `db:"to_user_id"`
	UnreadMessagesCount uint   `db:"unread_messages_count"`
	LastReadMessageID   uint64 `db:"last_read_message_id"`
}
