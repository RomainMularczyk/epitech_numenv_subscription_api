package repositories

import (
	"context"
	"fmt"
	"numenv_subscription_api/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"

	"github.com/google/uuid"
)

func Subscribe(ctx context.Context, user *models.Subscriber, sessionId string) error {
	client, err := db.Client()
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

	id, err := uuid.NewRandom()
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Error when generating UUID for a new subscriber.",
		)
		return err
	}
	uniqueStr, err := uuid.NewRandom()
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Error when generating unique string for a subscriber.",
		)
	}

	user.SetID(id.String())
	user.SetUniqueStr(uniqueStr.String())
	err = db.Exec[models.Subscriber](ctx, client, q, *user)
	if err != nil {
		return err
	}

	sessionInfos, err := GetSessionById(ctx, sessionId)
	if err != nil {
		return err
	}
	err = AddSubscriberToSession(ctx, sessionInfos.Id, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func ReadAll(ctx context.Context) ([]*models.Subscriber, error) {
	db, err := db.Client()
	if err != nil {
		fmt.Println("Error connecting to database", err)
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM subscribers")
	if err != nil {
		logs.Output(logs.ERROR, "Could not execute the query.")
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
		}
		result = append(result, &subscriber)
	}

	return result, nil
}
