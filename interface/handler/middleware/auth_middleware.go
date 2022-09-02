package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/infrastructure/jwter"
)

type AuthMiddleware struct {
	j *jwter.JWTer
}

func NewAuthMiddleware(j *jwter.JWTer) *AuthMiddleware {
	return &AuthMiddleware{j}
}

func (am *AuthMiddleware) JwtAuthenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.ParseRequest(c.Request(), jwt.WithKey(jwa.RS256, am.j.PublicKey))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		uid, err := am.j.Store.Load(c.Request().Context(), token.JwtID())
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		c = SetContext(c, uid)
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
