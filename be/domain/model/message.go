package model

import "time"

type MessageID uint64
type MessageType int

const (
	MessageTypeText  MessageType = 1
	MessageTypeImage MessageType = 2
)

type Message struct {
	ID             MessageID
	SendUserID     UserID
	ConversationID ConversationID
	Type           MessageType
	Text           string
	CreatedAt      time.Time
}

type Messages []*Message
