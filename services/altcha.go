package services

import (
	"math/big"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
	"numenv_subscription_api/services/altcha"
)

func Altcha() (*models.Challenge, error) {
  challenge, err := altcha.CreateALTCHA("", *big.NewInt(0))
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Error when creating Altcha challenge.",
    )
    return nil, err
  }

  return challenge, nil
}
