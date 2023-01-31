package repository

import (
	"context"

	"github.com/syougo1209/b-match-server/domain/model"
)

//go:generate mockgen -source=conversation_repository.go -destination=../../mock/repository/conversation.go
type ConversationRepository interface {
	UpdateLastMessageID(ctx context.Context, conversationID model.ConversationID, messageID model.MessageID) error
	FetchConversaionList(ctx context.Context, uid model.UserID) (model.Conversations, error)
}
