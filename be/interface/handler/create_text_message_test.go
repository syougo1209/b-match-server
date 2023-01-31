package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/interface/handler/middleware"
	"github.com/syougo1209/b-match-server/interface/presenter"
	mock_usecase "github.com/syougo1209/b-match-server/mock/usecase"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestCreateMessage_ServeHTTP(t *testing.T) {
	okJSON := `{"text":"test"}`
	invalidTextJSON := `{"text":""}`
	message := &model.Message{
		ID:             1,
		SendUserID:     1,
		ConversationID: 1,
		Type:           model.MessageTypeText,
		Text:           "test",
		CreatedAt:      time.Now(),
	}

	tests := map[string]struct {
		param         string
		prepareMockFn func(m *mock_usecase.MockCreateTextMessage)
		status        int
	}{
		"textが不適切な時": {
			param: invalidTextJSON,
			prepareMockFn: func(m *mock_usecase.MockCreateTextMessage) {
				m.EXPECT().Call(gomock.Any(), model.UserID(1), model.ConversationID(1), "test", gomock.Any()).Return(nil, errors.New("error")).AnyTimes()
			},
			status: http.StatusUnprocessableEntity,
		},
		"usecaseからerrNotFoundエラーが返ってくる時": {
			param: okJSON,
			prepareMockFn: func(m *mock_usecase.MockCreateTextMessage) {
				m.EXPECT().Call(gomock.Any(), model.UserID(1), model.ConversationID(1), "test", gomock.Any()).Return(nil, model.ErrNotFound)
			},
			status: http.StatusNotFound,
		},
		"usecaseからerrNotFound以外のエラーが返ってくる時": {
			param: okJSON,
			prepareMockFn: func(m *mock_usecase.MockCreateTextMessage) {
				m.EXPECT().Call(gomock.Any(), model.UserID(1), model.ConversationID(1), "test", gomock.Any()).Return(nil, errors.New("error"))
			},
			status: http.StatusInternalServerError,
		},
		"処理が適切に完了したとき": {
			param: okJSON,
			prepareMockFn: func(m *mock_usecase.MockCreateTextMessage) {
				m.EXPECT().Call(gomock.Any(), model.UserID(1), model.ConversationID(1), "test", gomock.Any()).Return(message, nil)
			},
			status: http.StatusCreated,
		},
	}
	for n, tt := range tests {

		t.Run(n, func(t *testing.T) {
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(tt.param))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)
			c.SetPath("/conversations/:id/messages")
			c.SetParamNames("id")
			c.SetParamValues("1")
			middleware.SetUserIDContext(c, 1)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_usecase.NewMockCreateTextMessage(ctrl)
			tt.prepareMockFn(mock)

			pre := presenter.MessagePresenter{}

			sut := CreateTextMessage{mock, pre, validator.New()}
			sut.ServeHTTP(c)
			resp := rec.Result()
			testutils.AssertResponseStatus(t,
				resp, tt.status,
			)
		})
	}
}
