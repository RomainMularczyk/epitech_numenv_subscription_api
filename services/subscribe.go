package services

import (
	"numenv_subscription_api/models"
	"numenv_subscription_api/repositories"

	"github.com/labstack/echo/v4"
)

func Subscribe(c echo.Context, u *models.Subscriber, id string) error {
	err := repositories.Subscribe(c.Request().Context(), u)
	if err != nil {
		return err
	}
	return nil
}

func ReadAll(c echo.Context) ([]*models.Subscriber, error) {
	result, err := repositories.ReadAll(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return result, nil
}
