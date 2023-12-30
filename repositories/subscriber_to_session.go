package repositories

import (
	"context"
	"fmt"
	"numenv_subscription_api/db"
	"numenv_subscription_api/errors/logs"

	"github.com/google/uuid"
)

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
	dbClient, err := db.Client()
	if err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Error when generating unique string for a subscriber.",
		)
	}

	q := `INSERT INTO subscribers_to_sessions (
    id, 
    sessions_id, 
    subscribers_id,
    unique_str
  ) VALUES ($1, $2, $3, $4)`

	_, err = dbClient.ExecContext(
    ctx, 
    q,
    id,
    sessionId,
    subscriberId,
    uniqueStr,
  )
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
