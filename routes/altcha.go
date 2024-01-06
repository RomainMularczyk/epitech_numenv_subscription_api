package routes

import (
  "numenv_subscription_api/controllers"
  "github.com/labstack/echo/v4"
)

func Altcha(e *echo.Echo) {
  e.GET("/altcha", controllers.Altcha)
}
