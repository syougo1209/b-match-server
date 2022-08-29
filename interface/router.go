package router

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/config"
)

func NewRouter(ctx context.Context, cfg *config.Config, xdb *sqlx.DB) (*echo.Echo, error) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world. you've requested\n")
	})
	return e, nil
}
