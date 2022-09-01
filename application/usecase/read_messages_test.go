package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/syougo1209/b-match-server/domain/model"
	mock_repository "github.com/syougo1209/b-match-server/mock/repository"
)

func TestReadMessages_Call(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		conversationID model.ConversationID
		messageID      model.MessageID
	}
	tests := map[string]struct {
		args          args
		prepareMockFn func(m *mock_repository.MockConversationStateRepository)
		wantErr       bool
	}{
		"repositoryがエラーを返すとき、エラーを返すこと": {
			args: args{model.ConversationID(1), model.MessageID(0)},
			prepareMockFn: func(m *mock_repository.MockConversationStateRepository) {
				m.EXPECT().ReadMessages(gomock.Any(), model.UserID(1), model.ConversationID(1), model.MessageID(0)).Return(errors.New("error"))
			},
			wantErr: true,
		},
		"正常に更新されたとき、エラーを返さないこと": {
			args: args{model.ConversationID(1), model.MessageID(0)},
			prepareMockFn: func(m *mock_repository.MockConversationStateRepository) {
				m.EXPECT().ReadMessages(gomock.Any(), model.UserID(1), model.ConversationID(1), model.MessageID(0)).Return(nil)
			},
			wantErr: false,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_repository.NewMockConversationStateRepository(ctrl)
			tt.prepareMockFn(mock)

			u := NewReadMessages(mock)
			err := u.Call(ctx, tt.args.conversationID, tt.args.messageID)
			if tt.wantErr == false && err != nil {
				t.Fatalf("want no error, but got %v", err)
				return
			}
		})
	}
}
