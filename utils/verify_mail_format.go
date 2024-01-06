package utils

import (
	"net/mail"
	"numenv_subscription_api/errors/logs"
)


func VerifyMailFormat(emailStr string) (error) {
  _, err := mail.ParseAddress(emailStr)
  if err != nil {
    logs.Output(
      logs.WARNING,
      "Could not parse user email.",
    )
    return err
  }
  return nil
}

