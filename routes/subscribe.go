package routes

import (
	"numenv_subscription_api/controllers"

	"github.com/labstack/echo/v4"
)

func Subscribe(e *echo.Echo) {
	e.POST("/subscribe/:id", controllers.Subscribe)
	e.GET("/subscribers", controllers.ReadAll)
}
