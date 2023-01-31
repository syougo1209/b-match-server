package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/interface/handler/middleware"
	"github.com/syougo1209/b-match-server/interface/presenter"
)

type FetchConversationList struct {
	UseCase   usecase.FetchConversationList
	Presenter presenter.ConversationPresenter
}

type fetchConversationListRequest struct {
}

func (fcl *FetchConversationList) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	uid, ok := middleware.GetUserIDContext(ctx)
	if !ok {
		log.Fatal("contextからuser idが取得できなかった")
		return c.JSON(http.StatusInternalServerError, "contextからuser idが取得できなかった")
	}
	conversations, err := fcl.UseCase.Call(ctx, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := fcl.Presenter.CreateConversationListRes(conversations)
	return c.JSON(http.StatusOK, res)
}
