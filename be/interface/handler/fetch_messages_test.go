package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/interface/presenter"
	mock_usecase "github.com/syougo1209/b-match-server/mock/usecase"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestFetchConversationMessages_ServeHTTP(t *testing.T) {
	t.Parallel()
	messages := model.Messages{
		{
			ID:             1,
			SendUserID:     1,
			ConversationID: 1,
			Type:           model.MessageTypeText,
			Text:           "test 1",
			CreatedAt:      time.Now(),
		},
		{
			ID:             2,
			SendUserID:     1,
			ConversationID: 1,
			Type:           model.MessageTypeImage,
			Text:           "test 2",
			CreatedAt:      time.Now(),
		},
	}
	type query struct {
		cursor, limit string
	}

	tests := map[string]struct {
		q             query
		prepareMockFn func(m *mock_usecase.MockFetchMessages)
		status        int
	}{
		"クエリパラメータがない場合、200が返ること": {
			q: query{"", ""},
			prepareMockFn: func(m *mock_usecase.MockFetchMessages) {
				m.EXPECT().Call(gomock.Any(), model.ConversationID(1), DefaultCursor, DefaultLimit).Return(messages, model.MessageID(0), nil)
			},
			status: http.StatusOK,
		},
		"クエリパラメータがある場合、200が返ること": {
			q: query{"2", "2"},
			prepareMockFn: func(m *mock_usecase.MockFetchMessages) {
				m.EXPECT().Call(gomock.Any(), model.ConversationID(1), 2, 2).Return(messages, model.MessageID(0), nil)
			},
			status: http.StatusOK,
		},
		"usecaseでエラーが発生する場合、500が返ること": {
			q: query{"1", "1"},
			prepareMockFn: func(m *mock_usecase.MockFetchMessages) {
				m.EXPECT().Call(gomock.Any(), model.ConversationID(1), 1, 1).Return(nil, model.MessageID(0), errors.New("error"))
			},
			status: http.StatusInternalServerError,
		},
		"クエリパラメータが不適切な値の場合、400が返ること": {
			q:             query{"invalid", "invalid"},
			prepareMockFn: func(m *mock_usecase.MockFetchMessages) {},
			status:        http.StatusBadRequest,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			e := echo.New()

			rec := httptest.NewRecorder()

			q := make(url.Values)
			q.Set("cursor", tt.q.cursor)
			q.Set("limit", tt.q.limit)
			req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

			c := e.NewContext(req, rec)
			c.SetPath("/conversations/:id/messages")
			c.SetParamNames("id")
			c.SetParamValues("1")

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_usecase.NewMockFetchMessages(ctrl)
			tt.prepareMockFn(mock)

			mp := presenter.MessagePresenter{}
			sut := FetchMessages{UseCase: mock, Presenter: mp}
			sut.ServeHTTP(c)

			resp := rec.Result()
			testutils.AssertResponseStatus(t,
				resp, tt.status,
			)
		})
	}
}
