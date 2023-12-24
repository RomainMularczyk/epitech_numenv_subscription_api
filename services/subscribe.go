package services

import (
	"numenv_subscription_api/models"
	"numenv_subscription_api/repositories"
	"numenv_subscription_api/services/mail"

	"github.com/labstack/echo/v4"
)

func Subscribe(c echo.Context, user *models.Subscriber, sessionName string) error {
	err := repositories.Subscribe(c.Request().Context(), user, sessionName)
	if err != nil {
		return err
	}
	mail.Send_mail()
	return nil
}

func ReadAll(c echo.Context) ([]*models.Subscriber, error) {
	result, err := repositories.ReadAll(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return result, nil
}
