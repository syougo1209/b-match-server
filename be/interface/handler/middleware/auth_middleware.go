package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/config"
	"github.com/syougo1209/b-match-server/domain/model"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

type AuthMiddleware struct {
	UseCase usecase.GetCurrentUserFromSub
}

func (am *AuthMiddleware) EnsureValidToken(cfg *config.Config) func(c echo.HandlerFunc) echo.HandlerFunc {
	issuerURL, err := url.Parse("https://" + cfg.Auth0Domain + "/")

	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{cfg.Auth0Audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "No Authorization Header")
			}
			if !strings.HasPrefix(authorization, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header")
			}
			token := strings.TrimPrefix(authorization, "Bearer ")
			claims, err := jwtValidator.ValidateToken(c.Request().Context(), token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
			}
			sub := claims.(*validator.ValidatedClaims).RegisteredClaims.Subject
			uid, err := am.UseCase.Call(c.Request().Context(), sub)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "failed to find user")
			}
			c = SetUserIDContext(c, *uid)
			return next(c)
		}
	}
}

type userIDKey struct{}

func SetUserIDContext(c echo.Context, uid model.UserID) echo.Context {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, userIDKey{}, uid)
	c.SetRequest(c.Request().WithContext(ctx))
	return c
}

func GetUserIDContext(ctx context.Context) (model.UserID, bool) {
	uid, ok := ctx.Value(userIDKey{}).(model.UserID)
	return uid, ok
}
