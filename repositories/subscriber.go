package repositories

import (
	"context"
	"fmt"
	"log"
	"numenv_subscription_api/db"
	"numenv_subscription_api/models"

	"github.com/google/uuid"
)

func Subscribe(ctx context.Context, user *models.Subscriber) error {
	db, err := db.Config()
	if err != nil {
		fmt.Println("Error connecting to database", err)
	}

	q := "INSERT INTO subscribers (id, first_name,last_name, email, unique_str, institution, epitech_degree) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	fmt.Println("q", q)
	insert, err := db.Prepare(q)
	if err != nil {
		fmt.Println("Error preparing statement", err)
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Error generating uuid", err)
		return err
	}

	_, err = insert.ExecContext(ctx,
		id,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.UniqueStr,
		user.Institution,
		user.EpitechDegree,
	)

	if err != nil {
		fmt.Println("Error inserting into database", err)
		return err
	}

	return nil
}

func ReadAll(ctx context.Context) ([]*models.Subscriber, error) {
	db, err := db.Config()
	if err != nil {
		fmt.Println("Error connecting to database", err)
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM subscribers")
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		result = append(result, &subscriber)
	}

	return result, nil
}
