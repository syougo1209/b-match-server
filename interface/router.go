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
	"github.com/syougo1209/b-match-server/interface/handler/middleware"
	"github.com/syougo1209/b-match-server/interface/presenter"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(ctx context.Context, cfg *config.Config, xdb *sqlx.DB) (*echo.Echo, error) {
	e := echo.New()
	v := validator.New()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	//repository
	mr := &database.MessageRepository{Db: xdb}
	csr := &database.ConversationStateRepository{Db: xdb}
	cr := &database.ConversationRepository{Db: xdb}
	tx := database.NewTransaction(xdb)
	//presenter
	mp := presenter.MessagePresenter{}
	cp := presenter.ConversationPresenter{}

	e.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world\n")
	})

	convesationGroup := e.Group("/conversations")
	ucfm := usecase.NewFetchMessages(mr)
	fmHandler := handler.FetchMessages{UseCase: ucfm, Presenter: mp}
	convesationGroup.GET("/:id/messages", fmHandler.ServeHTTP)

	ucrm := usecase.NewReadMessages(csr)
	rmHandler := handler.ReadMessages{UseCase: ucrm, Validator: v}
	convesationGroup.PATCH("/:id/read_message", rmHandler.ServeHTTP)

	ucctm := usecase.NewCreateMessage(mr, csr, cr, tx)
	ctmHandler := handler.CreateTextMessage{UseCase: ucctm, Presenter: mp, Validator: v}
	e.POST("/conversations/:id/messages", ctmHandler.ServeHTTP)

	meGroup := e.Group("/me")
	meGroup.Use(middleware.EnsureValidToken(cfg))
	ucfcl := usecase.NewFetchConversationList(cr)
	fclHandler := handler.FetchConversationList{UseCase: ucfcl, Presenter: cp}
	meGroup.GET("/conversations", fclHandler.ServeHTTP)

	return e, nil
}
