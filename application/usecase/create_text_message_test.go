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
	mock_usecase "github.com/syougo1209/b-match-server/mock/usecase"
)

func TestCreateTextMessage_Call(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	now := time.Now()
	message := &model.Message{
		ID:             1,
		SendUserID:     1,
		ConversationID: 1,
		Type:           model.MessageTypeText,
		Text:           "text",
		CreatedAt:      now,
	}
	tests := map[string]struct {
		prepareMessageMockFn      func(m *mock_repository.MockMessageRepository)
		prepareConversationMockFn func(m *mock_repository.MockConversationRepository)
		message                   *model.Message
		wantErr                   bool
	}{
		"message repositoryがエラーを返すとき, エラーを返すこと": {
			prepareMessageMockFn: func(m *mock_repository.MockMessageRepository) {
				m.EXPECT().CreateTextMessage(gomock.Any(), model.ConversationID(1), model.UserID(1), "text", now).Return(nil, errors.New("error"))
			},
			prepareConversationMockFn: func(m *mock_repository.MockConversationRepository) {
				m.EXPECT().UpdateLastMessageID(gomock.Any(), model.ConversationID(1), message.ID).Return(errors.New("error")).AnyTimes()
			},
			message: nil, wantErr: true,
		},
		"更新系の処理が失敗するとき、エラーを返すこと": {
			prepareMessageMockFn: func(m *mock_repository.MockMessageRepository) {
				m.EXPECT().CreateTextMessage(gomock.Any(), model.ConversationID(1), model.UserID(1), "text", now).Return(message, nil)
			},
			prepareConversationMockFn: func(m *mock_repository.MockConversationRepository) {
				m.EXPECT().UpdateLastMessageID(gomock.Any(), model.ConversationID(1), message.ID).Return(errors.New("error"))
			},
			message: nil, wantErr: true,
		},
		"正常に更新されたとき、エラーを返さないこと": {
			prepareMessageMockFn: func(m *mock_repository.MockMessageRepository) {
				m.EXPECT().CreateTextMessage(gomock.Any(), model.ConversationID(1), model.UserID(1), "text", now).Return(message, nil)
			},
			prepareConversationMockFn: func(m *mock_repository.MockConversationRepository) {
				m.EXPECT().UpdateLastMessageID(gomock.Any(), model.ConversationID(1), message.ID).Return(nil)
			},
			message: message, wantErr: false,
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			messageMock := mock_repository.NewMockMessageRepository(ctrl)
			conversationMock := mock_repository.NewMockConversationRepository(ctrl)
			conversationStateMock := mock_repository.NewMockConversationStateRepository(ctrl)
			conversationStateMock.EXPECT().IncrementMessageCount(gomock.Any(), model.UserID(1), model.ConversationID(1)).Return(nil).AnyTimes()
			tt.prepareMessageMockFn(messageMock)
			tt.prepareConversationMockFn(conversationMock)
			txMock := &mock_usecase.MockTx{}

			u := NewCreateMessage(messageMock, conversationStateMock, conversationMock, txMock)

			gotMessage, err := u.Call(ctx, model.UserID(1), model.ConversationID(1), "text", now)
			if tt.wantErr == false && err != nil {
				t.Fatalf("want no error, but got %v", err)
				return
			}
			if d := cmp.Diff(gotMessage, tt.message); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
