package utils

import (
	"numenv_subscription_api/errors/logs"
	"time"
)

func FormatDate(date string) time.Time {
  date, err := time.Parse(time.RFC3339, date)
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not parse the date.",
    )
  }

  return date.Format("DD-MM-YYYY")
}
