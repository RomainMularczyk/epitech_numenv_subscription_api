package routes

import (
	"numenv_subscription_api/controllers"
	"numenv_subscription_api/middlewares"

	"github.com/labstack/echo/v4"
)

func Subscribe(e *echo.Echo) {
	e.GET("/subscribe", controllers.ReadAll)
  // Groups
  g := e.Group("/subscribe")
  g.Use(middlewares.IsSessionFull)
	g.POST("/:speaker", controllers.Subscribe, middlewares.IsSessionFull)
}

