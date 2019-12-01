package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RegisterRouteBody struct {
	Url string `json:"url"`
}

type RegisterRouteResponse struct {
	Url string `json:"url"`
}

func (a *App) RegisterRoute(ctx echo.Context) error {
	body := new(RegisterRouteBody)
	err := ctx.Bind(body)
	if err != nil {
		return err
	}

	urlID, err := a.URLRepo.Create(ctx.Request().Context(), body.Url)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, RegisterRouteResponse{Url: urlID})
}

func (a *App) GetRoute(ctx echo.Context) error {
	id := ctx.Param("id")

	url, err := a.URLRepo.FindByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.Redirect(http.StatusPermanentRedirect, url.URL)
}
