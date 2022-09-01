package usecase

import (
	"context"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/domain/repository"
)

//go:generate mockgen -source=fetch_messages.go -destination=../../mock/usecase/fetch_messages.go
type FetchMessages interface {
	Call(ctx context.Context, conversationID model.ConversationID, cursor, limit int) (model.Messages, model.MessageID, error)
}

type fetchMessages struct {
	MessageRepo repository.MessageRepository
}

func NewFetchMessages(mr repository.MessageRepository) *fetchMessages {
	return &fetchMessages{
		MessageRepo: mr,
	}
}

func (fcm *fetchMessages) Call(
	ctx context.Context, conversationID model.ConversationID,
	cursor, limit int,
) (model.Messages, model.MessageID, error) {
	// next_cursorを返すために+1件多く取ってくる
	messages, err := fcm.MessageRepo.FetchMessages(
		ctx, conversationID,
		cursor, limit+1,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("FetchMessages: %w", err)
	}

	if len(messages) < limit+1 {
		return messages, 0, nil
	} else {
		next_cursor := messages[limit].ID
		return messages[0:limit], next_cursor, nil
	}
}
