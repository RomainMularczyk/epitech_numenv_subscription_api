package db

import (
  "numenv_subscription_api/errors/logs"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Client() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf(
    "host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		host,
    port,
    user,
    password,
    dbname,
  )

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
    logs.Output(
      logs.ERROR, 
      "Wrong database string connection format.",
    )
		return nil, err
	}

  err = db.Ping()
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not open connection to the database.",
    )
  }

	return db, nil
}
