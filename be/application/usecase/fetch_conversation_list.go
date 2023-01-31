package usecase

import (
	"context"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/domain/repository"
)

//go:generate mockgen -source=fetch_conversation_list.go -destination=../../mock/usecase/fetch_conversation_list.go
type FetchConversationList interface {
	Call(ctx context.Context, uid model.UserID) (model.Conversations, error)
}
type fetchConversationList struct {
	ConversationRepo repository.ConversationRepository
}

func NewFetchConversationList(cr repository.ConversationRepository) *fetchConversationList {
	return &fetchConversationList{
		ConversationRepo: cr,
	}
}

func (fcl *fetchConversationList) Call(ctx context.Context, uid model.UserID) (model.Conversations, error) {
	conversations, err := fcl.ConversationRepo.FetchConversaionList(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("FetchConversationList: %w", err)
	}
	return conversations, nil
}
