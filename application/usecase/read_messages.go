package usecase

import (
	"context"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/domain/repository"
)

//go:generate mockgen -source=read_messages.go -destination=../../mock/usecase/read_messages.go
type ReadMessages interface {
	Call(ctx context.Context, uid model.UserID, conversationID model.ConversationID, messageID model.MessageID) error
}
type readMessages struct {
	Repo repository.ConversationStateRepository
}

func NewReadMessages(csr repository.ConversationStateRepository) *readMessages {
	return &readMessages{
		Repo: csr,
	}
}

func (rm *readMessages) Call(ctx context.Context, uid model.UserID, conversationID model.ConversationID, messageID model.MessageID) error {
	err := rm.Repo.ReadMessages(ctx, uid, conversationID, messageID)
	if err != nil {
		return fmt.Errorf("readMessages Call: %w", err)
	}
	return nil
}
