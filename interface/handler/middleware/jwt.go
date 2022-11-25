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
			c = setSubContext(c, sub)
			return next(c)
		}
	}
}

type subKey struct{}

func setSubContext(c echo.Context, sub string) echo.Context {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, subKey{}, sub)
	c.SetRequest(c.Request().WithContext(ctx))
	return c
}

func getSubContext(ctx context.Context) (string, bool) {
	sub, ok := ctx.Value(subKey{}).(string)
	return sub, ok
}
