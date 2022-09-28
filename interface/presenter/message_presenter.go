package presenter

import (
	"time"

	"github.com/syougo1209/b-match-server/domain/model"
)

const (
	MessageTypeText  = "text"
	MessageTypeImage = "image"
)

type MessagePresenter struct{}

func (mp *MessagePresenter) CreateMessagesRes(messages model.Messages, nextCursor model.MessageID) MessagesWithPages {
	message_count := len(messages)
	response_messages := make([]MessageJSON, message_count)
	for i, v := range messages {
		mtype := convertType(v.Type)
		response_messages[i] = MessageJSON{
			ID:         uint64(v.ID),
			Content:    v.Text,
			Type:       mtype,
			SendUserId: uint64(v.SendUserID),
			CreatedAt:  v.CreatedAt,
		}
	}

	response := MessagesWithPages{
		Messages:   response_messages,
		NextCursor: uint64(nextCursor),
		Limit:      message_count,
	}
	return response
}

func (mp *MessagePresenter) CreateMessageRes(message model.Message) MessageJSON {
	mtype := convertType(message.Type)
	return MessageJSON{
		ID:         uint64(message.ID),
		Content:    message.Text,
		Type:       mtype,
		SendUserId: uint64(message.SendUserID),
		CreatedAt:  message.CreatedAt,
	}
}
func convertType(mtype model.MessageType) string {
	if mtype == model.MessageTypeText {
		return MessageTypeText
	} else {
		return MessageTypeImage
	}
}

type MessageJSON struct {
	ID         uint64    `json:"id"`
	Content    string    `json:"content"`
	Type       string    `json:"type"`
	SendUserId uint64    `json:"send_user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
type MessagesWithPages struct {
	Messages   []MessageJSON `json:"messages"`
	NextCursor uint64        `json:"next_cursor"`
	Limit      int           `json:"limit"`
}
