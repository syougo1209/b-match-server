package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/domain/model"
	mock_middleware "github.com/syougo1209/b-match-server/mock/middleware"
	"github.com/syougo1209/b-match-server/testutils"
)

func TestAuthMiddleware_JwtAuthenticate(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	uid := model.UserID(1)
	tests := map[string]struct {
		prepareMockFn func(m *mock_middleware.Mockjwter)
		status        int
	}{
		"認証に失敗するとき, 403を返すこと": {
			prepareMockFn: func(m *mock_middleware.Mockjwter) {
				m.EXPECT().CheckLoginState(gomock.Any(), req).Return(nil, errors.New("error"))
			},
			status: http.StatusForbidden,
		},
		"認証に成功するとき、後続の処理を続けること": {
			prepareMockFn: func(m *mock_middleware.Mockjwter) {
				m.EXPECT().CheckLoginState(gomock.Any(), req).Return(&uid, nil)
			},
			status: http.StatusOK,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			e := echo.New()
			rec := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_middleware.NewMockjwter(ctrl)
			tt.prepareMockFn(mock)

			m := NewAuthMiddleware(mock)
			c := e.NewContext(req, rec)
			h := m.JwtAuthenticate(echo.HandlerFunc(func(c echo.Context) error {
				return c.NoContent(http.StatusOK)
			}))

			h(c)
			resp := rec.Result()
			testutils.AssertResponseStatus(t,
				resp, tt.status,
			)
		})
	}
}
