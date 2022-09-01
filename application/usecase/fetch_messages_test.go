package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/syougo1209/b-match-server/domain/model"
	mock_repository "github.com/syougo1209/b-match-server/mock/repository"
)

func TestFetchMessages_Call(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	conversationID := model.ConversationID(1)
	messages := model.Messages{
		{
			ID:             1,
			SendUserID:     1,
			ConversationID: conversationID,
			Type:           model.MessageTypeText,
			Text:           "test 1",
			CreatedAt:      time.Now(),
		},
		{
			ID:             2,
			SendUserID:     1,
			ConversationID: conversationID,
			Type:           model.MessageTypeText,
			Text:           "test 2",
			CreatedAt:      time.Now(),
		},
	}

	type args struct {
		cursor, limit int
	}

	tests := map[string]struct {
		args          args
		prepareMockFn func(m *mock_repository.MockMessageRepository)
		wantMessages  model.Messages
		wantCursor    model.MessageID
		wantErr       bool
	}{
		"メッセージが全て取り切れていない場合にmessageIDとメッセージの配列が返ること": {
			args: args{2, 1},
			prepareMockFn: func(m *mock_repository.MockMessageRepository) {
				m.EXPECT().FetchMessages(gomock.Any(), conversationID, 2, 2).Return(messages, nil)
			},
			wantMessages: messages[0:1],
			wantCursor:   messages[1].ID,
			wantErr:      false,
		},
		"メッセージが全て取り切れている場合にmessageID=0とメッセージの配列が変えること": {
			args: args{2, 2},
			prepareMockFn: func(m *mock_repository.MockMessageRepository) {
				m.EXPECT().FetchMessages(gomock.Any(), conversationID, 2, 3).Return(messages, nil)
			},
			wantMessages: messages[0:2],
			wantCursor:   0,
			wantErr:      false,
		},
		"repositoryがエラーを返すとき、エラーを返すこと": {
			args: args{2, 2},
			prepareMockFn: func(m *mock_repository.MockMessageRepository) {
				m.EXPECT().FetchMessages(gomock.Any(), conversationID, 2, 3).Return(nil, errors.New("error"))
			},
			wantMessages: nil,
			wantCursor:   0,
			wantErr:      true,
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_repository.NewMockMessageRepository(ctrl)
			tt.prepareMockFn(mock)

			u := NewFetchMessages(mock)
			gotMessages, cursor, err := u.Call(ctx, conversationID, tt.args.cursor, tt.args.limit)
			if tt.wantErr == false && err != nil {
				t.Fatalf("want no error, but got %v", err)
				return
			}

			if d := cmp.Diff(gotMessages, tt.wantMessages); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
			if d := cmp.Diff(cursor, tt.wantCursor); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
