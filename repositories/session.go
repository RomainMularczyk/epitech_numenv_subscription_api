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

func GetSessionById(ctx context.Context, sessionId string) (*models.Session, error) {
  dbClient, err := db.Client()
  if err != nil {
    return nil, err
  }

  sess := &models.Session{}
  q := "SELECT id, name, speaker, date, type, num_subscribers FROM sessions WHERE id=$1"
  err = dbClient. 
    QueryRowContext(ctx, q, sessionId). 
    Scan(&sess.Id, &sess.Name, &sess.Speaker, &sess.Date, &sess.Type, &sess.NumSubscribers)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Error occurred when retrieving session by Id (%s). Query : %u\n",
        sessionId,
        err,
      ),
    )
    return nil, err
  }
  return sess, nil
}

func GetSessionBySpeaker(ctx context.Context, speaker string) (*models.Session, error) {
  dbClient, err := db.Client()
  if err != nil {
    return nil, err
  }

  sess := &models.Session{}
  q := "SELECT id, name, speaker, date, type, num_subscribers FROM sessions WHERE speaker=$1"
  err = dbClient.
    QueryRowContext(ctx, q, speaker).
    Scan(&sess.Id, &sess.Name, &sess.Speaker, &sess.Date, &sess.Type, &sess.NumSubscribers)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Error occurred when retrieving session by speaker name (%s). Query : %u\n",
        speaker,
        err,
      ),
    )
    return nil, err
  }
  return sess, nil
}

func GetSessionNumberSubscribers(
  ctx context.Context, 
  speaker string,
) (*int, error) {
  dbClient, err := db.Client()
  if err != nil {
    return nil, err
  }

  var count int
  err = dbClient.QueryRow("SELECT COUNT(*) FROM subscribers_to_sessions").Scan(&count)
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not query the of rows.",
    )
    return nil, err
  }

  return &count, nil
}

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
