package controllers

import (
	"fmt"
	"net/http"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
	"numenv_subscription_api/models/responses"
	"numenv_subscription_api/services"
  "numenv_subscription_api/utils"

	"github.com/labstack/echo/v4"
)

func Subscribe(ctx echo.Context) error {
	subscriber := &models.SubscriberWithChallenge{}
	err := ctx.Bind(subscriber)
	speaker := ctx.Param("speaker")

  // Verify user email format
  err = utils.VerifyMailFormat(subscriber.Email)
  if err != nil { return err }

  // Verify user metadata format
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

	err = services.SubscribeToSessionAndSendMail(
    ctx.Request().Context(), 
    subscriber, 
    speaker,
  )
	if err != nil {
		return ctx.JSON(
			http.StatusUnprocessableEntity,
			responses.ErrorResponse{
				Message: fmt.Sprintf(
					"Service could not handle the request. Error: %s.",
					err.Error(),
				),
			},
		)
	}

  subscriberWithoutChallenge := models.FilterOutAltcha(subscriber)

	return ctx.JSON(
		http.StatusAccepted,
		responses.SuccessResponse[*models.Subscriber]{
			Data:    &subscriberWithoutChallenge,
			Message: "Successfully registered new subscriber.",
		},
	)
}

// Retrieve all subscribers
func GetAllSubscribers(ctx echo.Context) error {
	subcribersList, err := services.GetAllSubscribers(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(
    http.StatusOK,
    responses.SuccessResponse[[]*models.Subscriber] {
      Data: subcribersList,
      Message: "Successfully retrieved all subscribers.",
    },
  )
}
