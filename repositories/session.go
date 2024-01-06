package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"numenv_subscription_api/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
)

// Get the session metadata by the user unique string
func GetSessionByUniqueStr(uniqueStr string) (*models.Session, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	sess := &models.Session{}
	q := `SELECT 
    sessions.id, 
    name, 
    speaker, 
    date, 
    type,
    discord_role_id,
    num_subscribers
    FROM sessions 
    JOIN subscribers_to_sessions ON sessions.id = subscribers_to_sessions.sessions_id 
    WHERE subscribers_to_sessions.unique_str=$1`

	stmt, err := client.Prepare(q)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Could not prepare query. Query : %s.",
				q,
			),
		)
		return nil, err
	}

	err = stmt.
		QueryRow(uniqueStr).
		Scan(
			&sess.Id,
			&sess.Name,
			&sess.Speaker,
			&sess.Date,
			&sess.Type,
			&sess.DiscordRoleId,
			&sess.NumSubscribers,
		)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Error occurred when retrieving session by UniqueStr (%s). Query : %u\n",
				uniqueStr,
				err,
			),
		)
		return nil, err
	}
	return sess, nil
}

// Get the session metadata by session ID
func GetSessionById(ctx context.Context, sessionId string) (*models.Session, error) {
	dbClient, err := db.Client()
	if err != nil {
		return nil, err
	}

	sess := &models.Session{}
	q := `SELECT 
    id, 
    name, 
    speaker, 
    date, 
    type, 
    discord_role_id, 
    num_subscribers 
    FROM sessions WHERE id=$1`

	stmt, err := dbClient.Prepare(q)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Could not prepare query. Query : %s.",
				q,
			),
		)
		return nil, err
	}

	err = stmt.
		QueryRowContext(ctx, sessionId).
		Scan(
			&sess.Id,
			&sess.Name,
			&sess.Speaker,
			&sess.Date,
			&sess.Type,
			&sess.DiscordRoleId,
			&sess.NumSubscribers,
		)
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

// Get the session metadata by speaker name
func GetSessionBySpeaker(
	ctx context.Context,
	speaker string,
) (*models.Session, error) {
	dbClient, err := db.Client()
	if err != nil {
		return nil, err
	}

	sess := &models.Session{}
	q := `SELECT 
    id, 
    name, 
    speaker, 
    date, 
    type, 
    discord_role_id, 
    num_subscribers 
    FROM sessions WHERE speaker=$1`

	stmt, err := dbClient.Prepare(q)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Could not prepare query. Query : %s.",
				q,
			),
		)
		return nil, err
	}

	err = stmt.
		QueryRowContext(ctx, speaker).
		Scan(
			&sess.Id,
			&sess.Name,
			&sess.Speaker,
			&sess.Date,
			&sess.Type,
			&sess.DiscordRoleId,
			&sess.NumSubscribers,
		)
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

// Count the number of subscribers for a given session
// queried by speaker name
func GetSessionNumberSubscribersBySpeaker(
	ctx echo.Context,
) (*int, error) {
	dbClient, err := db.Client()
	if err != nil {
		return nil, err
	}

	speaker := ctx.Param("speaker")

	var count int
	q := `SELECT COUNT(*) FROM subscribers_to_sessions 
    JOIN sessions ON sessions.id = subscribers_to_sessions.sessions_id
    WHERE sessions.speaker=$1`

	stmt, err := dbClient.Prepare(q)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Could not prepare query. Query : %s.",
				q,
			),
		)
		return nil, err
	}

	err = stmt.QueryRow(speaker).Scan(&count)
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Could not get number of subscribers from database.",
		)
		return nil, err
	}

	return &count, nil
}

// Get the session metadata by session name
func GetSessionByName(ctx context.Context, name string) (*models.Session, error) {
	dbClient, err := db.Client()
	if err != nil {
		return nil, err
	}

	sess := &models.Session{}
	q := "SELECT id, name, num_subscribers FROM sessions WHERE name=$1"

	stmt, err := dbClient.Prepare(q)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Could not prepare query. Query : %s.",
				q,
			),
		)
		return nil, err
	}

	err = stmt.
		QueryRowContext(ctx, name).
		Scan(&sess.Id, &sess.Name, &sess.NumSubscribers)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Error occurred with Get session by name (%s) query: %v\n",
				name,
				err,
			),
		)
		if errors.Is(err, sql.ErrNoRows) {
			logs.Output(logs.WARNING, "no sessions found with that ID")
			return nil, err
		}
		return nil, err
	}

	return sess, nil
}

func GetAllSessions() ([]*models.Session, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	q := `SELECT 
    id, 
    name, 
    speaker, 
    type, 
    date, 
    num_subscribers, 
    discord_role_id 
    FROM sessions`

	stmt, err := client.Prepare(q)
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Could not prepare query. Query : %s.",
				q,
			),
		)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		logs.Output(
			logs.ERROR,
			fmt.Sprintf(
				"Could not execute the query. Query : %s, produced error : %s.",
				q,
				err,
			),
		)
	}

	var result []*models.Session
	for rows.Next() {
		var session models.Session

		err := rows.Scan(
			&session.Id,
			&session.Name,
			&session.Speaker,
			&session.Type,
			&session.Date,
			&session.NumSubscribers,
			&session.DiscordRoleId,
		)
		if err != nil {
			logs.Output(
				logs.ERROR,
				"Values retrieved from database did not match model properties.",
			)
			return nil, err
		}
		result = append(result, &session)
	}

	return result, nil
}
