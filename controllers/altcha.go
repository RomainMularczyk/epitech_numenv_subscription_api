package controllers

import (
	"net/http"
	"numenv_subscription_api/models/responses"
	"numenv_subscription_api/services"

	"github.com/labstack/echo/v4"
)

func Altcha(ctx echo.Context) error {
  challenge, err := services.Altcha()
  if err != nil {
    return ctx.JSON(
      http.StatusInternalServerError,
      responses.ErrorResponse {
        Message: "Error happened when creating Altcha challenge.",
      },
    )
  }

  return ctx.JSON(
    http.StatusOK,
    challenge,
  )
}
