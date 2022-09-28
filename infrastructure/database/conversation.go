package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
)

type Conversation struct {
	ID            uint64 `db:"id"`
	LastMessageID uint64 `db:"last_message_id"`
}

type ConversationRepository struct {
	Db DbConnection
}

func (cr *ConversationRepository) UpdateLastMessageID(ctx context.Context, conversationID model.ConversationID, messageID model.MessageID) error {
	var dto Conversation
	query := `
	  select *
	  from conversation
		WHERE id =?
	`
	if err := cr.Db.GetContext(ctx, &dto, query, conversationID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("failed to GetContext conversation by id=%d: %w", conversationID, model.ErrNotFound)
		}
		return fmt.Errorf("failed to GetContext conversation by id=%d: %w", conversationID, model.ErrNotFound)
	}

	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to start transaction")
	}
	uquery := `
		UPDATE
		conversation
		SET last_message_id=?
		WHERE id =?
	`
	if _, err := tx.ExecContext(ctx, uquery, messageID, conversationID); err != nil {
		return fmt.Errorf("failed to update conversationState: %w", err)
	}
	return nil
}
