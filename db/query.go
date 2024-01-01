package db

import (
	"context"
	"database/sql"
	"numenv_subscription_api/errors/logs"
	"reflect"
)

func Exec[T any](
	ctx context.Context,
	client *sql.DB,
	q string,
	data T,
) error {
	query, err := client.Prepare(q)
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Error occurred with the query.",
		)
		return err
	}

	numFields := reflect.ValueOf(data).Type().NumField()
	arrayValues := make([]interface{}, numFields)
	for i := 0; i < numFields; i++ {
		value := reflect.ValueOf(data).Field(i).Interface()
		arrayValues[i] = value
	}

	_, err = query.ExecContext(
		ctx,
		arrayValues...,
	)
  defer query.Close()
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Error happened with the following SQL query: " + q,
		)
		return err
	}

	return nil
}
