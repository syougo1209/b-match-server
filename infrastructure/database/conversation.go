package database

type Conversation struct {
	ID            uint64 `db:"id"`
	LastMessageID uint64 `db:"last_message_id"`
}
