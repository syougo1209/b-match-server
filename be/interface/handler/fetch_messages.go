package handler

import (
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/interface/presenter"
)

var (
	DefaultLimit  = 15
	DefaultCursor = math.MaxInt
)

type FetchMessages struct {
	UseCase   usecase.FetchMessages
	Presenter presenter.MessagePresenter
}
type fetchMessagesRequest struct {
	ID     int `param:"id"`
	Cursor int `query:"cursor"`
	Limit  int `query:"limit"`
}

func (fcm *FetchMessages) ServeHTTP(c echo.Context) error {
	req := new(fetchMessagesRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	limit := req.Limit
	if limit == 0 {
		limit = DefaultLimit
	}
	cursor := req.Cursor
	if cursor == 0 {
		cursor = DefaultCursor
	}

	ctx := c.Request().Context()
	messages, nextCursor, err := fcm.UseCase.Call(ctx, model.ConversationID(req.ID), cursor, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	presenter := presenter.MessagePresenter{}
	res := presenter.CreateMessagesRes(messages, nextCursor)
	return c.JSON(http.StatusOK, res)
}
