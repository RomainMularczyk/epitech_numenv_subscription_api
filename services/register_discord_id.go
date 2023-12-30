package services

import (
  "numenv_subscription_api/models"
	"numenv_subscription_api/repositories"
)

func RegisterDiscordId(discordId string, uniqueStr string) (*models.Session, error) {
  repositories.RegisterUserDiscordId(discordId, uniqueStr)

  sess, err := repositories.GetSessionByUniqueStr(uniqueStr)
  if err != nil {
    return nil, err
  }

  return sess, nil
}
