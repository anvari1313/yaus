package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserRegisterBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	InsertedID string `json:"inserted_id"`
	Token      string `json:"token"`
}

func (a *App) UserRegisterHandler(c echo.Context) error {
	body := new(UserRegisterBody)
	err := c.Bind(body)
	if err != nil {
		return err
	}

	objectID, err := a.UserRepo.Create(c.Request().Context(), body.Name, body.Username, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UserRegisterResponse{InsertedID: objectID, Token: "some"})
}
