package repositories

import (
	"context"
	"fmt"
	"numenv_subscription_api/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"

	"github.com/google/uuid"
)

// Register a user in the database and create the intermediate
// table entry to associate it with the corresponding session
func Subscribe(
  ctx context.Context,
  subscriber *models.Subscriber,
  sessionId string,
) error {
	client, err := db.Client()
	if err != nil {
    return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	q := `INSERT INTO subscribers (
    id, 
    first_name,
    last_name,
    email,
    institution,
    epitech_degree,
    discord_id
  ) VALUES ($1, $2, $3, $4, $5, $6, NULLIF($7, ''));`

	subscriber.SetID(id.String())
	err = db.Exec[models.Subscriber](ctx, client, q, *subscriber)
	if err != nil {
		return err
	}

	return nil
}

// Register a user's Discord Id to the database
func RegisterSubscriberDiscordId(
  discordId string,
  uniqueStr string,
) error {
  client, err := db.Client()
  if err != nil {
    return err
  }
  defer client.Close()

  q := `UPDATE subscribers
    SET discord_id=$1
    FROM subscribers_to_sessions
    WHERE subscribers.id=subscribers_to_sessions.subscribers_id
    AND subscribers_to_sessions.unique_str=$2`
  stmt, err := client.Prepare(q)
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not prepare the statement.",
    )
    return err
  }

  _, err = stmt.Exec(discordId, uniqueStr)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not execute the query : %s.", q,
      ),
    )
    return err
  }

  return nil
}

// Retrieve user metadata by user email
func GetSubscriberByEmail(email string) (*models.Subscriber, error) {
  client, err := db.Client()
  if err != nil {
    return nil, err
  }
  defer client.Close()

  subscriber := &models.Subscriber{}
  q := `SELECT
    id,
    first_name,
    last_name,
    email,
    institution,
    epitech_degree,
    discord_id
    FROM subscribers WHERE email=$1
  `
  stmt, err := client.Prepare(q)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not prepare the statement : %s.",
        q,
      ),
    )
    return nil, err
  }

  err = stmt.QueryRow(email).Scan(
    &subscriber.Id,
    &subscriber.Firstname,
    &subscriber.Lastname,
    &subscriber.Email,
    &subscriber.Institution,
    &subscriber.EpitechDegree,
    &subscriber.DiscordId,
  ) 
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not execute the query: %s. Error produced : %s.",
        q,
        err,
      ),
    )
    return nil, err
  }

  return subscriber, nil
}

// Retrieve user metadata by user Id
func GetSubscriberById(id string) (*models.Subscriber, error) {
	client, err := db.Client()
	if err != nil {
    return nil, err
	}
  defer client.Close()

  subscriber := &models.Subscriber{}
  q := `SELECT 
      id, first_name, last_name, email, institution, epitech_degree, discord_id
      FROM subscribers WHERE id=$1`

  stmt, err := client.Prepare(q)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not prepare the statement : %s.",
        q,
      ),
    )
    return nil, err
  }

  err = stmt.QueryRow(id).Scan(
    &subscriber.Id,
    &subscriber.Firstname,
    &subscriber.Lastname,
    &subscriber.Email,
    &subscriber.Institution,
    &subscriber.EpitechDegree,
    &subscriber.DiscordId,
  )
	if err != nil {
		logs.Output(
      logs.ERROR, 
      fmt.Sprintf(
        "Could not execute the query: %s. Error procuded : %s.",
        q,
        err,
      ),
    )
    return nil, err
	}

  return subscriber, nil
}

func GetSubscriberByDiscordId(discordId string) (*models.Subscriber, error) {
  client, err := db.Client()
  if err != nil {
    return nil, err
  }
  defer client.Close()

  subscriber := &models.Subscriber{}
  q := `SELECT
    id, 
    first_name, 
    last_name, 
    email, 
    institution, 
    epitech_degree, 
    discord_id
    FROM subscribers
    WHERE discord_id=$1
  `

  stmt, err := client.Prepare(q)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not prepare the query. Query : %s.",
        q,
      ),
    )
    return nil, err
  }

  err = stmt.QueryRow(discordId).Scan(
    &subscriber.Id,
    &subscriber.Firstname,
    &subscriber.Lastname,
    &subscriber.Email,
    &subscriber.Institution,
    &subscriber.EpitechDegree,
    &subscriber.DiscordId,
  )
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not execute the query : %s, produced error : %s.",
        q,
        err,
      ),
    )
    return nil, err
  }

  return subscriber, nil
}

// Read all entries in the subscribers table
func GetAllSubscribers(ctx context.Context) ([]*models.Subscriber, error) {
	client, err := db.Client()
	if err != nil {
    return nil, err
	}
	defer client.Close()

  q := "SELECT * FROM subscribers"

  stmt, err := client.Prepare(q)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not prepare the query. Query : %s.",
        q,
      ),
    )
    return nil, err
  }

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		logs.Output(
      logs.ERROR, 
      fmt.Sprintf(
        "Could not execute the query : %s, produced error : %s.",
        q,
        err,
      ),
    )
    return nil, err
	}

	var subscribers []*models.Subscriber
	for rows.Next() {
		var subscriber models.Subscriber

		err := rows.Scan(
			&subscriber.Id,
			&subscriber.Firstname,
			&subscriber.Lastname,
			&subscriber.DiscordId,
			&subscriber.Email,
			&subscriber.Institution,
			&subscriber.EpitechDegree,
    )
		if err != nil {
			logs.Output(
				logs.ERROR,
				"Values retrieved from database did not match model properties.",
			)
      return nil, err
		}
		subscribers = append(subscribers, &subscriber)
	}

	return subscribers, nil
}
