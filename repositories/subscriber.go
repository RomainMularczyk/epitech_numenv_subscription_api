package repositories

import (
	"context"
	"numenv_subscription_api/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"

	"github.com/google/uuid"
)

// Register a user in the database and create the intermediate 
// table entry to associate it with the corresponding session
func Subscribe(
  ctx context.Context,
  user *models.Subscriber,
  sessionId string,
  uniqueStr string,
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
    discord_id,
    unique_str
  ) VALUES ($1, $2, $3, $4, $5, $6, NULLIF($7, ''), $8);`

	user.SetID(id.String())
	user.SetUniqueStr(uniqueStr)
	err = db.Exec[models.Subscriber](ctx, client, q, *user)
	if err != nil {
		return err
	}

	sessionInfos, err := GetSessionById(ctx, sessionId)
	if err != nil {
		return err
	}

	err = AddSubscriberToSession(
    ctx,
    sessionInfos.Id,
    user.Id,
    uniqueStr,
  )
	if err != nil {
		return err
	}
	return nil
}

// Register a user's Discord Id to the database
func RegisterUserDiscordId(
  uniqueStr string,
  discordId string,
) error {
  client, err := db.Client()
  if err != nil {
    return err
  }

  q := `UPDATE subscribers
    SET discord_id = $1
    WHERE unique_str = $2
  `

  query, err := client.Prepare(q)
  if err != nil {
    return err
  }

  _, err = query.Exec(q)

  return nil
}


// Read all entries in the subscribers table
func ReadAll(ctx context.Context) ([]*models.Subscriber, error) {
	db, err := db.Client()
	if err != nil {
    return nil, err
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM subscribers")
	if err != nil {
		logs.Output(logs.ERROR, "Could not execute the query.")
    return nil, err
	}
	defer rows.Close()

	var result []*models.Subscriber
	for rows.Next() {
		var subscriber models.Subscriber

		err := rows.Scan(
			&subscriber.Id,
			&subscriber.Firstname,
			&subscriber.Lastname,
			&subscriber.DiscordId,
			&subscriber.Email,
			&subscriber.UniqueStr,
			&subscriber.Institution,
			&subscriber.EpitechDegree)
		if err != nil {
			logs.Output(
				logs.ERROR,
				"Values retrieved from database did not match model properties.",
			)
      return nil, err
		}
		result = append(result, &subscriber)
	}

	return result, nil
}
