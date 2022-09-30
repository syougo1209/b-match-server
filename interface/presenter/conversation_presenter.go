package presenter

import (
	"time"

	"github.com/syougo1209/b-match-server/domain/model"
)

var (
	LastMessageImageText = "画像が送信されました"
)

type ConversationPresenter struct{}

func (cp *ConversationPresenter) CreateConversationListRes(conversations model.Conversations) ConversationsResponse {
	res := make([]ConversationJSON, len(conversations))
	for i, conversation := range conversations {
		toUser := userResponse{
			ID:   uint64(conversation.ToUser.ID),
			Name: conversation.ToUser.Name,
		}

		var lastMessageText string
		if conversation.LastMessage.Type == model.MessageTypeText {
			lastMessageText = conversation.LastMessage.Text
		} else {
			lastMessageText = LastMessageImageText
		}

		res[i] = ConversationJSON{
			ID:                  uint64(conversation.ID),
			UnreadMessagesCount: uint(conversation.UnreadMessagesCount),
			UpdatedAt:           conversation.LastMessage.CreatedAt,
			LastMessage:         lastMessageText,
			ToUser:              toUser,
		}
	}

	return ConversationsResponse{Conversations: res}
}

type ConversationJSON struct {
	ID                  uint64       `json:"id"`
	UnreadMessagesCount uint         `json:"unread_message_count"`
	UpdatedAt           time.Time    `json:"updated_at"`
	LastMessage         string       `json:"last_message"`
	ToUser              userResponse `json:"to_user"`
}

type ConversationsResponse struct {
	Conversations []ConversationJSON `json:"conversations"`
}
