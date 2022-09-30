package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/domain/repository"
)

//go:generate mockgen -source=create_text_message.go -destination=../../mock/usecase/create_text_message.go
type CreateTextMessage interface {
	Call(ctx context.Context, cid model.ConversationID, text string, now time.Time) (*model.Message, error)
}
type createMessage struct {
	messageRepo          repository.MessageRepository
	convesationStateRepo repository.ConversationStateRepository
	convesationRepo      repository.ConversationRepository
	transaction          Transaction
}

func NewCreateMessage(
	mr repository.MessageRepository, csr repository.ConversationStateRepository,
	cr repository.ConversationRepository, tx Transaction,
) *createMessage {
	return &createMessage{
		messageRepo:          mr,
		convesationStateRepo: csr,
		convesationRepo:      cr,
		transaction:          tx,
	}
}

func (cm *createMessage) Call(ctx context.Context, cid model.ConversationID, text string, now time.Time) (*model.Message, error) {
	m, err := cm.transaction.BeginTx(ctx, func(ctx context.Context) (interface{}, error) {
		return cm.createMessage(ctx, cid, text, now)
	})
	if m != nil {
		return m.(*model.Message), err
	} else {
		return nil, err
	}
}

func (cm *createMessage) createMessage(ctx context.Context, cid model.ConversationID, text string, now time.Time) (*model.Message, error) {
	uid := model.UserID(1)

	m, err := cm.messageRepo.CreateTextMessage(ctx, cid, uid, text, now)
	if err != nil {
		return nil, fmt.Errorf("messageRepo Create: %w", err)
	}
	//todo メッセージ詳細に相手がいるときはincrementせずにlast_read_messsage_idを更新したい
	err = cm.convesationStateRepo.IncrementMessageCount(ctx, uid, cid)
	if err != nil {
		return nil, fmt.Errorf("ConversationStateRepo Update: %w", err)
	}

	err = cm.convesationRepo.UpdateLastMessageID(ctx, cid, m.ID)
	if err != nil {
		return nil, fmt.Errorf("ConversationRepo Update: %w", err)
	}

	return m, nil
}
