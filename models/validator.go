package models

import (
	"github.com/go-playground/validator"
	"numenv_subscription_api/errors/logs"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		logs.Output(
			logs.ERROR,
			"Could not parse user metadata.",
		)
		return err
	}
	return nil
}
