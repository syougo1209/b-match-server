// Code generated by MockGen. DO NOT EDIT.
// Source: message_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/syougo1209/b-match-server/domain/model"
)

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// CreateTextMessage mocks base method.
func (m *MockMessageRepository) CreateTextMessage(ctx context.Context, conversationID model.ConversationID, uid model.UserID, text string, now time.Time) (*model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTextMessage", ctx, conversationID, uid, text, now)
	ret0, _ := ret[0].(*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTextMessage indicates an expected call of CreateTextMessage.
func (mr *MockMessageRepositoryMockRecorder) CreateTextMessage(ctx, conversationID, uid, text, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTextMessage", reflect.TypeOf((*MockMessageRepository)(nil).CreateTextMessage), ctx, conversationID, uid, text, now)
}

// FetchMessages mocks base method.
func (m *MockMessageRepository) FetchMessages(ctx context.Context, conversationID model.ConversationID, cursor, limit int) (model.Messages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMessages", ctx, conversationID, cursor, limit)
	ret0, _ := ret[0].(model.Messages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessages indicates an expected call of FetchMessages.
func (mr *MockMessageRepositoryMockRecorder) FetchMessages(ctx, conversationID, cursor, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessages", reflect.TypeOf((*MockMessageRepository)(nil).FetchMessages), ctx, conversationID, cursor, limit)
}