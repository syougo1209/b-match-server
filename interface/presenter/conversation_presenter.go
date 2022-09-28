package presenter

import (
	"time"

	"github.com/syougo1209/b-match-server/domain/model"
)

type ConversationPresenter struct{}

func (cp *ConversationPresenter) CreateConversationListRes(conversations []model.Conversation) ConversationsResponse {
	res := make(ConversationsResponse, len(conversations))
	for i, conversation := range conversations {
		toUser := userResponse{
			ID:   uint64(conversation.ToUser.ID),
			Name: conversation.ToUser.Name,
		}
		res[i] = ConversationJSON{
			ID:                  uint64(conversation.ID),
			UnreadMessagesCount: uint(conversation.UnreadMessagesCount),
			UpdatedAt:           conversation.LastMessage.CreatedAt,
			ToUser:              toUser,
		}
	}
	return res
}

type ConversationJSON struct {
	ID                  uint64       `json:"id"`
	UnreadMessagesCount uint         `json:"unread_message_count"`
	UpdatedAt           time.Time    `json:"updated_at"`
	LastMessage         string       `json:"last_message"`
	ToUser              userResponse `json:"to_user"`
}

type ConversationsResponse []ConversationJSON
