package repository

import (
	"context"

	"github.com/syougo1209/b-match-server/domain/model"
)

//go:generate mockgen -source=message_repository.go -destination=../../mock/repository/message.go
type MessageRepository interface {
	FetchMessages(ctx context.Context, conversationID model.ConversationID, cursor, limit int) (model.Messages, error)
}
