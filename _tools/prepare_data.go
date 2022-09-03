package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/b-match-server/domain/model"
)

func main() {
	ctx := context.Background()
	mysqlDB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/b-match?parseTime=true")

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer mysqlDB.Close()

	xdb := sqlx.NewDb(mysqlDB, "mysql")

	user := prepareUser(ctx, xdb)
	otherUser := prepareUser(ctx, xdb)
	cid := prepareConversation(ctx, xdb)
	prepareConversationState(ctx, xdb, cid, user, otherUser)
	prepareConversationState(ctx, xdb, cid, otherUser, user)
	prepareMessages(ctx, xdb, cid, user.ID)
	prepareMessages(ctx, xdb, cid, otherUser.ID)
}

func prepareUser(ctx context.Context, db *sqlx.DB) model.User {
	u := &model.User{
		Name:      "example",
		CreatedAt: time.Now(),
	}
	result, _ := db.ExecContext(ctx, "INSERT INTO user (name, created_at) VALUES (?, ?)",
		u.Name, u.CreatedAt)

	id, _ := result.LastInsertId()

	u.ID = model.UserID(id)
	return *u
}
func prepareConversation(ctx context.Context, db *sqlx.DB) model.ConversationID {
	result, _ := db.ExecContext(ctx, "INSERT INTO conversation (last_message_id) VALUES (?)", 0)
	id, _ := result.LastInsertId()
	return model.ConversationID(id)
}

func prepareConversationState(
	ctx context.Context, db *sqlx.DB,
	cid model.ConversationID, fromUser model.User, toUser model.User,
) *model.Conversation {
	c := model.Conversation{
		ID:                  cid,
		FromUser:            fromUser,
		ToUser:              toUser,
		UnreadMessagesCount: 0,
		LastReadMessage:     nil,
		LastMessage:         nil,
	}
	db.ExecContext(
		ctx,
		"INSERT INTO conversation_state (conversation_id, from_user_id, to_user_id, unread_messages_count, last_read_message_id) VALUES (?, ?, ?, ?, ?)",
		c.ID, c.FromUser.ID, c.ToUser.ID,
		c.UnreadMessagesCount, 0,
	)
	return &c
}
func prepareMessages(
	ctx context.Context, db *sqlx.DB,
	cid model.ConversationID, uid model.UserID,
) {
	messages := model.Messages{
		{
			SendUserID:     uid,
			ConversationID: cid,
			Type:           1,
			Text:           "初めまして",
			CreatedAt:      time.Now(),
		},
		{
			SendUserID:     uid,
			ConversationID: cid,
			Type:           1,
			Text:           "こんにちは",
			CreatedAt:      time.Now(),
		},
		{
			SendUserID:     uid,
			ConversationID: cid,
			Type:           1,
			Text:           "よろしくお願いします",
			CreatedAt:      time.Now(),
		},
	}
	db.ExecContext(ctx,
		`INSERT INTO message (send_user_id, conversation_id, type, text, created_at)
			VALUES
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?);`,
		messages[0].SendUserID, messages[0].ConversationID, messages[0].Type, messages[0].Text, messages[0].CreatedAt,
		messages[1].SendUserID, messages[1].ConversationID, messages[1].Type, messages[1].Text, messages[1].CreatedAt,
		messages[2].SendUserID, messages[2].ConversationID, messages[2].Type, messages[2].Text, messages[2].CreatedAt,
	)
}
