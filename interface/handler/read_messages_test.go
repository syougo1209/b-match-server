package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/domain/model"
	mock_usecase "github.com/syougo1209/b-match-server/mock/usecase"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestReadMessages_ServeHTTP(t *testing.T) {
	t.Parallel()

	okJSON := `{"read_message_id":1}`
	invalidJSON := `{"read_message_id":""}`
	type param struct {
		ID   string
		JSON string
	}
	tests := map[string]struct {
		param         param
		prepareMockFn func(m *mock_usecase.MockReadMessages)
		status        int
	}{
		"パラメータが不適切な時": {
			param: param{"1", invalidJSON},
			prepareMockFn: func(m *mock_usecase.MockReadMessages) {
				m.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
			status: http.StatusBadRequest},
		"usecaseからerrNotFoundエラーが返ってくる時": {
			param: param{"1", okJSON},
			prepareMockFn: func(m *mock_usecase.MockReadMessages) {
				m.EXPECT().Call(gomock.Any(), model.ConversationID(1), model.MessageID(1)).Return(model.ErrNotFound)
			},
			status: http.StatusNotFound,
		},
		"usecaseからerrNotFound以外のエラーが返ってくる時": {
			param: param{"1", okJSON},
			prepareMockFn: func(m *mock_usecase.MockReadMessages) {
				m.EXPECT().Call(gomock.Any(), model.ConversationID(1), model.MessageID(1)).Return(errors.New("error"))
			},
			status: http.StatusInternalServerError,
		},
		"処理が適切に完了したとき": {
			param: param{"1", okJSON},
			prepareMockFn: func(m *mock_usecase.MockReadMessages) {
				m.EXPECT().Call(gomock.Any(), model.ConversationID(1), model.MessageID(1)).Return(nil)
			},
			status: http.StatusOK,
		},
	}

	for n, tt := range tests {

		t.Run(n, func(t *testing.T) {
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(tt.param.JSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)
			c.SetPath("/conversations/:id/read_message")
			c.SetParamNames("id")
			c.SetParamValues(tt.param.ID)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_usecase.NewMockReadMessages(ctrl)
			tt.prepareMockFn(mock)

			sut := ReadMessages{mock, validator.New()}
			sut.ServeHTTP(c)
			resp := rec.Result()
			testutils.AssertResponseStatus(t,
				resp, tt.status,
			)
		})
	}
}
