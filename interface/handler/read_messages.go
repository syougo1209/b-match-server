package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/domain/model"
)

type ReadMessages struct {
	UseCase usecase.ReadMessages
}
type readMessagesRequest struct {
	ID        uint64 `param:"id"`
	MessageID uint64 `param:"message_id"`
}

func (rm *ReadMessages) ServeHTTP(c echo.Context) error {
	req := new(readMessagesRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err := rm.UseCase.Call(ctx, model.ConversationID(req.ID), model.MessageID(req.MessageID))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
