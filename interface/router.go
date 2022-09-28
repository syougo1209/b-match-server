package router

import (
	"context"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/config"
	"github.com/syougo1209/b-match-server/infrastructure/database"
	"github.com/syougo1209/b-match-server/infrastructure/jwter"
	"github.com/syougo1209/b-match-server/infrastructure/redis"
	"github.com/syougo1209/b-match-server/interface/handler"
	"github.com/syougo1209/b-match-server/interface/handler/middleware"
	"github.com/syougo1209/b-match-server/interface/presenter"
)

func NewRouter(ctx context.Context, cfg *config.Config, xdb *sqlx.DB) (*echo.Echo, error) {
	e := echo.New()
	v := validator.New()

	rcli, err := redis.NewRedis(ctx, cfg)
	if err != nil {
		log.Printf("failed to start redis: %v", err)
		return nil, err
	}
	jwter, err := jwter.NewJWTer(rcli)
	if err != nil {
		log.Printf("failed to start jwter: %v", err)
		return nil, err
	}

	//repository
	mr := &database.MessageRepository{Db: xdb}
	csr := &database.ConversationStateRepository{Db: xdb}
	cr := &database.ConversationRepository{Db: xdb}
	ur := &database.UserRepository{Db: xdb}
	tx := database.NewTransaction(xdb)
	//presenter
	mp := presenter.MessagePresenter{}

	authMiddleware := middleware.NewAuthMiddleware(jwter)

	ucel := usecase.NewEasyLogin(ur, jwter)
	elHandler := handler.EasyLogin{UseCase: ucel}
	e.POST("/login/easy", elHandler.ServeHTTP)

	e.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world\n")
	})

	convesationGroup := e.Group("/conversations")
	convesationGroup.Use(authMiddleware.JwtAuthenticate)
	ucfm := usecase.NewFetchMessages(mr)
	fmHandler := handler.FetchMessages{UseCase: ucfm, Presenter: mp}
	convesationGroup.GET("/:id/messages", fmHandler.ServeHTTP)

	ucrm := usecase.NewReadMessages(csr)
	rmHandler := handler.ReadMessages{UseCase: ucrm, Validator: v}
	convesationGroup.PATCH("/:id/read_message", rmHandler.ServeHTTP)

	ucctm := usecase.NewCreateMessage(mr, csr, cr, tx)
	ctmHandler := handler.CreateTextMessage{UseCase: ucctm, Presenter: mp, Validator: v}
	e.POST("/conversations/:id/messages", ctmHandler.ServeHTTP)

	return e, nil
}
