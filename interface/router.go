package router

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/config"
	"github.com/syougo1209/b-match-server/infrastructure/database"
	"github.com/syougo1209/b-match-server/interface/handler"
	"github.com/syougo1209/b-match-server/interface/presenter"
)

func NewRouter(ctx context.Context, cfg *config.Config, xdb *sqlx.DB) (*echo.Echo, error) {
	e := echo.New()
	v := validator.New()

	//repository
	mr := &database.MessageRepository{Db: xdb}
	csr := &database.ConversationStateRepository{Db: xdb}
	//presenter
	mp := presenter.MessagePresenter{}

	ucfm := usecase.NewFetchMessages(mr)
	fmHandler := handler.FetchMessages{UseCase: ucfm, Presenter: mp}
	e.GET("/conversations/:id/messages", fmHandler.ServeHTTP)

	ucrm := usecase.NewReadMessages(csr)
	rmHandler := handler.ReadMessages{UseCase: ucrm, Validator: v}
	e.PATCH("/conversations/:id/read_message", rmHandler.ServeHTTP)

	e.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world\n")
	})
	return e, nil
}
