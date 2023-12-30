package services

import (
	"fmt"
	dbError "numenv_subscription_api/errors/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
	"numenv_subscription_api/repositories"
	"numenv_subscription_api/services/mail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// Register a user to the database and creates a unique string
// that is used to also register the same user on Discord so
// it can access the corresponding channel
func Subscribe(
  c echo.Context, 
  subscriber *models.Subscriber, 
  speaker string,
) error {
  // Retrieve session metadata
  sess, err := repositories.GetSessionBySpeaker(
    c.Request().Context(), 
    speaker,
  )

  // Generate unique string
  uniqueStr, err := uuid.NewRandom()
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Error when generating unique string for a subscriber.",
    )
    return err
  }

  // Register subscriber in subscriber table
	err = repositories.Subscribe(
    c.Request().Context(), 
    subscriber,
    sess.Id,
  )
  // If the email used is already registered, it will only
  // add the subscriber Id in the intermediate table
  if err != nil {
    if !dbError.IsErrorCode(err, pq.ErrorCode("23505")) {
      return err
    } else {
      oldSubscriber, err := repositories.GetSubscriberByEmail(subscriber.Email)
      if err != nil {
        return err
      }
      subscriber.SetID(oldSubscriber.Id)
    }
  }

  // /!\ Handle case if error with transactions

  // Register suscriber in intermediate table
  err = repositories.AddSubscriberToSession(
    c.Request().Context(),
    sess.Id,
    subscriber.Id,
    uniqueStr.String(),
  )
  if err != nil { return err }

  mail.SendMail(subscriber.Email, sess.Name, uniqueStr.String())

	return nil
}

// Get user metadata by user Id
func GetSubscriberByUniqueStr(uniqueStr string) (*models.Subscriber, error) {
  subscriberId, err := repositories.
    GetSubscriberForeignKeyByUniqueStr(uniqueStr)
  if err != nil {
    return nil, err
  }

  subscriber, err := repositories.GetSubscriberById(*subscriberId)
  if err != nil {
    return nil, err
  }

  return subscriber, nil
}

// Get all the subscribers
func ReadAll(c echo.Context) ([]*models.Subscriber, error) {
	result, err := repositories.ReadAll(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return result, nil
}
