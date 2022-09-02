package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syougo1209/b-match-server/application/usecase"
	"github.com/syougo1209/b-match-server/domain/model"
)

type EasyLogin struct {
	UseCase usecase.EasyLogin
}
type easyLoginRequest struct {
	ID uint64 `json:"uid" validate:"required"`
}

func (el *EasyLogin) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(easyLoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	jwt, err := el.UseCase.Call(ctx, model.UserID(req.ID))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: jwt,
	}

	return c.JSON(http.StatusOK, res)
}
