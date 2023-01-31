package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/interface/handler/middleware"
	"github.com/syougo1209/b-match-server/interface/presenter"
)

type CreateTextMessage struct {
	UseCase   usecase.CreateTextMessage
	Presenter presenter.MessagePresenter
	Validator *validator.Validate
}

type requestParam struct {
	ConversationID uint64 `param:"id"`
	Text           string `json:"text" validate:"min=1,lt=1024,required"`
}

func (cm *CreateTextMessage) ServeHTTP(c echo.Context) error {
	req := new(requestParam)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if err := cm.Validator.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	ctx := c.Request().Context()
	now := time.Now()
	uid, ok := middleware.GetUserIDContext(ctx)
	if !ok {
		log.Fatal("contextからuser idが取得できなかった")
		return c.JSON(http.StatusInternalServerError, "contextからuser idが取得できなかった")
	}
	message, err := cm.UseCase.Call(ctx, uid, model.ConversationID(req.ConversationID), req.Text, now)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			log.Printf("Could not access a resource that should have existed: %v", err)
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := cm.Presenter.CreateMessageRes(*message)
	return c.JSON(http.StatusCreated, res)
}
