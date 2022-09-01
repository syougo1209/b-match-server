package model

type ConversationID uint64

type Conversation struct {
	ID          ConversationID
	LastMessage *Message
}
