package repository

import (
	"context"
	"time"

	"github.com/syougo1209/b-match-server/domain/model"
)

//go:generate mockgen -source=message_repository.go -destination=../../mock/repository/message.go
type MessageRepository interface {
	FetchMessages(ctx context.Context, conversationID model.ConversationID, cursor, limit int) (model.Messages, error)
	CreateTextMessage(ctx context.Context, conversationID model.ConversationID, uid model.UserID, text string, now time.Time) (*model.Message, error)
}
