package controllers

import (
  "github.com/labstack/echo/v4"
  "numenv_subscription_api/services"
)

func Subscribe(ctx echo.Context) error {
  // do stuff
  services.Subscribe()
  return nil
}
