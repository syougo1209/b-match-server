package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/interface/presenter"
)

type FetchConversationList struct {
	UseCase   usecase.FetchConversations
	Presenter presenter.ConversationPresenter
}

type fetchConversationListRequest struct {
}

func (fcl *FetchConversationList) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	conversations, err := fcl.UseCase.Call(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := fcl.Presenter.CreateConversationListRes(*conversations)
	return c.JSON(http.StatusOK, res)
}
