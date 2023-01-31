package mock_usecase

import context "context"

type MockTx struct {
}

func (m *MockTx) BeginTx(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
	return f(ctx)
}
