package services

import (
	"numenv_subscription_api/models"
	"numenv_subscription_api/repositories"
  "numenv_subscription_api/errors/logs"
	"numenv_subscription_api/services/mail"

  "github.com/labstack/echo/v4"
  "github.com/google/uuid"
)

// Register a user to the database and creates a unique string 
// that is used to also register the same user on Discord so 
// it can access the corresponding channel
func Subscribe(
  c echo.Context, 
  user *models.Subscriber, 
  speaker string,
) error {
  sess, err := repositories.GetSessionBySpeaker(
    c.Request().Context(), 
    speaker,
  )

  uniqueStr, err := uuid.NewRandom()
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Error when generating unique string for a subscriber.",
    )
  }

	err = repositories.Subscribe(
    c.Request().Context(), 
    user,
    sess.Id,
    uniqueStr.String(),
  )
	if err != nil {
		return err
	}

  mail.SendMail(user.Email, sess.Name, uniqueStr.String())

	return nil
}

// Get all the subscribers
func ReadAll(c echo.Context) ([]*models.Subscriber, error) {
	result, err := repositories.ReadAll(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return result, nil
}
