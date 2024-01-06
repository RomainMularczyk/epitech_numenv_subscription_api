package mail

import (
	"fmt"
	"numenv_subscription_api/errors/logs"
	"os"

	. "github.com/mailjet/mailjet-apiv3-go"
)

func SendMail(address string, session string, uniqueStr string) {
	mailjetClient := NewMailjetClient(
      os.Getenv("MAILJET_API_KEY"),
      os.Getenv("MAILJET_API_SECRET"),
    )
	email := &InfoSendMail {
      FromEmail: os.Getenv("MAILJET_SENDER_ADDRESS"),
      FromName: "Team .env{2024}",
      Subject: fmt.Sprintf(".env{2024} - Inscription %s", session),
      HTMLPart: FormatContent(session, uniqueStr),
      Recipients: []Recipient {
        Recipient {
          Email: address,
        },
      },
    }
  
	_, err := mailjetClient.SendMail(email)
	if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not send the email.",
    )
	}
}
