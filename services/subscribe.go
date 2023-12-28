package services

import (
	"numenv_subscription_api/models"
	"numenv_subscription_api/repositories"
  "numenv_subscription_api/errors/logs"
	"numenv_subscription_api/services/mail"
  "numenv_subscription_api/services/discord"

	"github.com/labstack/echo/v4"
  "github.com/google/uuid"
)

func Subscribe(
  c echo.Context, 
  user *models.Subscriber, 
  speaker string,
) error {
  discord.DiscordClient()
  sess, err := repositories.GetSessionBySpeaker(c.Request().Context(), speaker)

	err = repositories.Subscribe(c.Request().Context(), user, sess.Id)
	if err != nil {
		return err
	}

  uniqueStr, err := uuid.NewRandom()
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Error when generating unique string for a subscriber.",
    )
  }

  mail.SendMail(user.Email, sess.Name, uniqueStr.String())
	return nil
}

func ReadAll(c echo.Context) ([]*models.Subscriber, error) {
	result, err := repositories.ReadAll(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return result, nil
}
