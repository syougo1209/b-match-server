package model

import "time"

type ConversationID uint64

type LastMessage struct {
	Type      MessageType
	Text      string
	CreatedAt time.Time
}
type Conversation struct {
	ID                  ConversationID
	ToUser              User
	FromUser            *User
	UnreadMessagesCount uint
	LastMessage         *LastMessage
}

type Conversations []*Conversation
