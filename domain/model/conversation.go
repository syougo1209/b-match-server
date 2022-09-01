package model

type ConversationID uint64

type Conversation struct {
	ID                  ConversationID
	ToUser              User
	FromUser            User
	UnreadMessagesCount uint
	LastReadMessage     *Message
	LastMessage         *Message
}
