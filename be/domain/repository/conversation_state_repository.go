package repository

import (
	"context"

	"github.com/syougo1209/b-match-server/domain/model"
)

//go:generate mockgen -source=conversation_state_repository.go -destination=../../mock/repository/conversation_state.go
type ConversationStateRepository interface {
	ReadMessages(ctx context.Context, uid model.UserID, conversationID model.ConversationID, messageID model.MessageID) error
	IncrementMessageCount(ctx context.Context, toUID model.UserID, conversationID model.ConversationID) error
}
