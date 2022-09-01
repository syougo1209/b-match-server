package usecase

import (
	"context"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/domain/repository"
)

type ReadMessages interface {
	Call(ctx context.Context, conversationID model.ConversationID, messageID model.MessageID) error
}
type readMessages struct {
	Repo repository.ConversationStateRepository
}

func NewReadMessages(csr repository.ConversationStateRepository) *readMessages {
	return &readMessages{
		Repo: csr,
	}
}

func (rm *readMessages) Call(ctx context.Context, conversationID model.ConversationID, messageID model.MessageID) error {
	//uidをauthから取得する
	uid := model.UserID(1)
	err := rm.Repo.ReadMessages(ctx, uid, conversationID, messageID)
	if err != nil {
		return fmt.Errorf("readMessages Call: %w", err)
	}
	return nil
}
