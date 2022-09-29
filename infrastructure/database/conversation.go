package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/b-match-server/domain/model"
)

type Conversation struct {
	ID            uint64 `db:"id"`
	LastMessageID uint64 `db:"last_message_id"`
}

type ConversationInfo struct {
	ConversationID       uint64    `db:"conversation_id"`
	ToUserID             uint64    `db:"to_user_id"`
	UnreadMessagesCount  uint      `db:"unread_messages_count"`
	LastMessageID        uint64    `db:"last_message_id"`
	LastMessageType      uint      `db:"type"`
	LastMessageText      string    `db:"text"`
	LastMessageCreatedAt time.Time `db:"created_at"`
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

func (cr *ConversationRepository) FetchConversaionList(ctx context.Context, uid model.UserID) (model.Conversations, error) {
	conversationsDTO := []*ConversationInfo{}
	cquery := `
		select cs.conversation_id, cs.to_user_id, cs.unread_messages_count, c.last_message_id, m.type, m.text, m.created_at
		from conversation as c
		join conversation_state as cs on c.id = cs.conversation_id
		join message as m on m.id = c.last_message_id
		WHERE cs.from_user_id = ?
		order by m.created_at desc
	`
	if err := cr.Db.SelectContext(ctx, conversationsDTO, cquery, uid); err != nil {
		return nil, fmt.Errorf("failed to SelectContext conversationInfo by from_user_id=%d: %w", uid, err)
	}

	toUserIDs := make([]uint64, len(conversationsDTO))
	for i, v := range conversationsDTO {
		toUserIDs[i] = v.ToUserID
	}
	toUsers := []User{}
	uquery := `
    select *
		from user
		where id in (?)
	`
	usql, uparams, err := sqlx.In(uquery, toUserIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to create In query: %w", err)
	}
	userQuery := cr.Db.Rebind(usql)

	if err := cr.Db.SelectContext(ctx, toUsers, userQuery, uparams...); err != nil {
		return nil, fmt.Errorf("failed to SelectContext toUsers: %w", err)
	}

	conversations := make(model.Conversations, len(conversationsDTO))
	for ci, cv := range conversationsDTO {
		conversations[ci] = &model.Conversation{
			ID:                  model.ConversationID(cv.ConversationID),
			UnreadMessagesCount: cv.UnreadMessagesCount,
			LastMessage: &model.LastMessage{
				Type:      model.MessageType(cv.LastMessageType),
				Text:      cv.LastMessageText,
				CreatedAt: cv.LastMessageCreatedAt,
			},
		}
		for _, uv := range toUsers {
			if cv.ToUserID == uv.ID {
				conversations[ci].ToUser = model.User{
					ID:        model.UserID(uv.ID),
					Name:      uv.Name,
					CreatedAt: uv.CreatedAt,
				}
			}
		}
	}
	return conversations, nil
}
