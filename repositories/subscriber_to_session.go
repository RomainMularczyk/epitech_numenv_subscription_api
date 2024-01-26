package repositories

import (
	"context"
	"fmt"
	"numenv_subscription_api/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"

	"github.com/google/uuid"
)

// Retrieve all sessions a subscriber is registered to
func GetAllSessionsBySubscriberId(
  id string,
) ([]*models.Session, error) {
  client, err := db.Client()  
  if err != nil {
    return nil, err
  }
  defer client.Close()

  q := `SELECT
    sessions.id,
    name,
    speaker,
    date,
    type,
    discord_role_id,
    num_subscribers
    FROM sessions
    JOIN subscribers_to_sessions
    ON sessions.id=subscribers_to_sessions.sessions_id
    WHERE subscribers_to_sessions.subscribers_id=$1
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

  rows, err := stmt.Query(id)
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

  var sessions []*models.Session
  for rows.Next() {
    var session models.Session

    err := rows.Scan(
      &session.Id,
      &session.Name,
      &session.Speaker,
      &session.Date,
      &session.Type,
      &session.DiscordRoleId,
      &session.NumSubscribers,
    )
    if err != nil {
      logs.Output(
        logs.ERROR,
        "Values retrieved from database did not match model properties.",
      )
      return nil, err
    }
    sessions = append(sessions, &session)
  }

  return sessions, nil
}

// Retrieve user by querying intermediate table
func GetSubscriberForeignKeyByUniqueStr(
  uniqueStr string,
) (*string, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
  defer client.Close()

  var subscriberId string
  q := `SELECT subscribers_id 
    FROM subscribers_to_sessions 
    WHERE unique_str=$1`

  stmt, err := client.Prepare(q)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not prepare the query : %s.",
        q,
      ),
    )
    return nil, err
  }

  err = stmt.
    QueryRow(uniqueStr).
    Scan(&subscriberId)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Error occurred when retrieving subscriber by UniqueStr (%s). Query : %s\n",
        uniqueStr,
        q,
      ),
    )
    return nil, err
  }

  return &subscriberId, nil
}

// Register data in the intermediate table
// Those should include :
// - The session ID
// - The subscriber ID
// - A unique string which is also sent by email to the subscriber
func AddSubscriberToSession(
  ctx context.Context, 
  sessionId string, 
  subscriberId string,
  uniqueStr string,
) error {
	client, err := db.Client()
	if err != nil {
		return err
	}
  defer client.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Error when generating unique string for a subscriber.",
		)
    return err
	}

	q := `INSERT INTO subscribers_to_sessions (
    id, 
    sessions_id, 
    subscribers_id,
    unique_str
  ) VALUES ($1, $2, $3, $4)`

  stmt, err := client.Prepare(q)
  if err != nil {
    logs.Output(
      logs.ERROR,
      fmt.Sprintf(
        "Could not prepare the query : %s.",
        q,
      ),
    )
    return err
  }

  _, err = stmt.ExecContext(
    ctx, 
    id,
    sessionId,
    subscriberId,
    uniqueStr,
  )
	if err != nil {
		logs.Output(logs.ERROR, fmt.Sprintf(
			`Error occured with Add subscriber (id: %s) to session (id: %s).
      Query: %s\n, produced error : %s`,
			subscriberId,
			sessionId,
			q,
      err,
		))
		return err
	}

	return nil
}
