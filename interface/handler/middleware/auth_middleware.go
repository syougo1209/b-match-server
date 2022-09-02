package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/domain/model"
)

type AuthMiddleware struct {
	j jwter
}
type jwter interface {
	CheckLoginState(context.Context, *http.Request) (*model.UserID, error)
}

func NewAuthMiddleware(j jwter) *AuthMiddleware {
	return &AuthMiddleware{j}
}

func (am *AuthMiddleware) JwtAuthenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid, err := am.j.CheckLoginState(c.Request().Context(), c.Request())
		if err != nil {
			return c.JSON(http.StatusForbidden, err.Error())
		}
		c = SetContext(c, *uid)
		return next(c)
	}
}

type userIDKey struct{}

func SetContext(c echo.Context, uid model.UserID) echo.Context {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, userIDKey{}, uid)
	c.SetRequest(c.Request().WithContext(ctx))
	return c
}
