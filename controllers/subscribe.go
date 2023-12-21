package controllers

import (
	"errors"
	"net/http"
	"numenv_subscription_api/models"
	"numenv_subscription_api/services"

	"github.com/labstack/echo/v4"
)

func Subscribe(ctx echo.Context) error {
	u := &models.Subscriber{}
	err := ctx.Bind(u)
	id := ctx.Param("id")

	if err != nil {
		return errors.New("invalid user query: " + err.Error())
	}
	err = services.Subscribe(ctx, u, id)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func ReadAll(ctx echo.Context) error {
	list, err := services.ReadAll(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, list)
}
