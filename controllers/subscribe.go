package controllers

import (
	"fmt"
	"net/http"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
	"numenv_subscription_api/models/responses"
	"numenv_subscription_api/services"

	"github.com/labstack/echo/v4"
)

func Subscribe(ctx echo.Context) error {
	user := &models.Subscriber{}
	err := ctx.Bind(user)
	sessionName := ctx.Param("speaker")

	if err != nil {
		logs.Output(
			logs.ERROR,
			"Could not parse user metadata.",
		)
		return ctx.JSON(
			http.StatusUnprocessableEntity,
			responses.ErrorResponse{Message: "Could not parse user metadata."},
		)
	}
	err = services.SubscribeToSessionAndSendMail(ctx.Request().Context(), user, sessionName)
	if err != nil {
		return ctx.JSON(
			http.StatusUnprocessableEntity,
			responses.ErrorResponse{
				Message: fmt.Sprintf(
					"Service could not handle the request. Error: %s",
					err.Error(),
				),
			},
		)
	}
	return ctx.JSON(
		http.StatusCreated,
		responses.SuccessResponse[models.Subscriber]{
			Data:    *user,
			Message: "Successfully registered new subscriber.",
		},
	)
}

// Retrieve all subscribers
func GetAllSubscribers(ctx echo.Context) error {
	list, err := services.GetAllSubscribers(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, list)
}
