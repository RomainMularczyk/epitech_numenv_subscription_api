package routes

import (
	"github.com/labstack/echo/v4"
	"numenv_subscription_api/controllers"
)

func Subscribe(e *echo.Echo) {
  e.POST("/subscribe/:id", controllers.Subscribe) 
}
