package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"numenv_subscription_api/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
)

func GetSessionByName(ctx context.Context, name string) (*models.Session, error) {
	dbClient, err := db.Client()
	if err != nil {
		return nil, err
	}

	sess := &models.Session{}
	q := "SELECT id, name, num_subscribers FROM sessions WHERE name=$1"
	err = dbClient.QueryRowContext(ctx, q, name).Scan(&sess.Id, &sess.Name, &sess.NumSubscribers)
	if err != nil {
		logs.Output(logs.ERROR, fmt.Sprintf("error occured with Get session by name (%s) query: %v\n", name, err))
		if errors.Is(err, sql.ErrNoRows) {
			logs.Output(logs.WARNING, "no sessions found with that ID")
			return nil, err
		}
		return nil, err
	}

	return sess, nil
}

func AddSubscriberToSession(ctx context.Context, sessionId string, subscriberId string) error {
	dbClient, err := db.Client()
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

	q := "INSERT INTO subscribers_to_sessions (id, sessions_id, subscribers_id) VALUES ($1, $2, $3)"
	_, err = dbClient.ExecContext(ctx, q, uniqueStr, sessionId, subscriberId)
	if err != nil {
		logs.Output(logs.ERROR, fmt.Sprintf(
			"Error occured with Add subscriber (id: %s) to session (id: %s) Query: %v\n",
			sessionId,
			subscriberId,
			err,
		))
		return err
	}

	return nil
}
