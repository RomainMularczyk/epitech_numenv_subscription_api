package routes

import (
	"github.com/go-playground/validator"
	"numenv_subscription_api/controllers"
	"numenv_subscription_api/middlewares"
	"numenv_subscription_api/models"

	"github.com/labstack/echo/v4"
)

func Subscribe(e *echo.Echo) {
	// Groups
	e.Validator = &models.CustomValidator{Validator: validator.New()}
	g := e.Group("/subscribe")
	g.Use(middlewares.IsSessionFull)
	g.POST("/:speaker", controllers.Subscribe, middlewares.IsSessionFull)
}
