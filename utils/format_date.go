package utils

import (
	"fmt"
	"numenv_subscription_api/errors/logs"
	"time"
)

func monthInFrench(m time.Month) string {
	mois := map[time.Month]string{
		time.January:   "Janvier",
		time.February:  "Février",
		time.March:     "Mars",
		time.April:     "Avril",
		time.May:       "Mai",
		time.June:      "Juin",
		time.July:      "Juillet",
		time.August:    "Août",
		time.September: "Septembre",
		time.October:   "Octobre",
		time.November:  "Novembre",
		time.December:  "Décembre",
	}

	return mois[m]
}

func FormatDate(date string) string {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Could not parse the date.",
		)
	}
	parsedDate = parsedDate.In(time.Local)
	formattedDate := fmt.Sprintf("%v %s %v, à %s",
		parsedDate.Day(),
		monthInFrench(parsedDate.Month()),
		parsedDate.Year(),
		parsedDate.Format("15:04"))
	return formattedDate
}
