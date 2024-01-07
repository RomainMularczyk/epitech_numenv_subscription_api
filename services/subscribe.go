package services

import (
	"context"
	altchaError "numenv_subscription_api/errors/altcha"
	dbError "numenv_subscription_api/errors/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
	"numenv_subscription_api/repositories"
	"numenv_subscription_api/services/altcha"
	"numenv_subscription_api/services/mail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// Register a subscriber to the database and creates a unique string
// that is used to also register the same subscriber on Discord so
// it can access the corresponding channel
func SubscribeToSession(
	ctx context.Context,
	subscriber *models.Subscriber,
	speaker string,
) (*models.Session, *string, error) {
	// Retrieve session metadata
	sess, err := repositories.GetSessionBySpeaker(
		ctx,
		speaker,
	)
	if err != nil {
		return nil, nil, err
	}

	// Generate unique string
	uuid, err := uuid.NewRandom()
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Error when generating unique string for a subscriber.",
		)
		return nil, nil, err
	}

	// Register subscriber in subscriber table
	err = repositories.Subscribe(
		ctx,
		subscriber,
		sess.Id,
	)

	// If the email used is already registered, it will only
	// add the subscriber Id in the intermediate table
	if err != nil {
		if !dbError.IsErrorCode(err, pq.ErrorCode("23505")) {
			return nil, nil, err
		} else {
			oldSubscriber, err := repositories.GetSubscriberByEmail(subscriber.Email)
			if err != nil {
				return nil, nil, err
			}
			subscriber.SetID(oldSubscriber.Id)
		}
	}

	// Register subscriber in intermediate table
	uniqueStr := uuid.String()
	err = repositories.AddSubscriberToSession(
		ctx,
		sess.Id,
		subscriber.Id,
		uniqueStr,
	)
	if err != nil {
		if !dbError.IsErrorCode(err, pq.ErrorCode("23505")) {
			return nil, nil, err
		} else {
			return nil, nil, dbError.AlreadyRegisteredError{
				Message: "Subscriber is already registered to this session",
			}
		}
	}

	return sess, &uniqueStr, nil
}

// Register a subscriber to a session
// Send the mail with the unique string
func SubscribeToSessionAndSendMail(
	ctx context.Context,
	subscriber *models.SubscriberWithChallenge,
	speaker string,
) error {

	// Verify Altcha
	res, err := altcha.VerifyALTCHA(*subscriber)
	if err != nil {
		return err
	}
	if !res {
		return altchaError.AltchaNotMatchingError{
			Message: "Altcha challenge could not be validated.",
		}
	}

	newSubscriber := models.FilterOutAltcha(subscriber)

	sess, uniqueStr, err := SubscribeToSession(ctx, &newSubscriber, speaker)
	if err != nil {
		return err
	}

	mail.SendMail(subscriber.Email, sess.Name, *uniqueStr)

	return nil
}

// Get subscriber metadata by subscriber Id
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
func GetAllSubscribers(c echo.Context) ([]*models.Subscriber, error) {
	result, err := repositories.GetAllSubscribers(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return result, nil
}
