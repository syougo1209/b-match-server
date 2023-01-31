package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/interface/handler/middleware"
)

type ReadMessages struct {
	UseCase   usecase.ReadMessages
	Validator *validator.Validate
}
type readMessagesRequest struct {
	ID            uint64 `param:"id"`
	ReadMessageID uint64 `json:"read_message_id" validate:"required"`
}

func (rm *ReadMessages) ServeHTTP(c echo.Context) error {
	req := new(readMessagesRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := rm.Validator.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	uid, ok := middleware.GetUserIDContext(ctx)
	if !ok {
		log.Fatal("contextからuser idが取得できなかった")
		return c.JSON(http.StatusInternalServerError, "contextからuser idが取得できなかった")
	}
	err := rm.UseCase.Call(ctx, uid, model.ConversationID(req.ID), model.MessageID(req.ReadMessageID))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
