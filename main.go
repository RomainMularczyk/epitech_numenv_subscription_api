package main

import (
	"net/http"
	"numenv_subscription_api/routes"
	"numenv_subscription_api/services/discord"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  go func() {
    discord.DiscordClient()
  }()

  // Starting server
	e := echo.New()
  // CORS configuration
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
    AllowOrigins: []string{os.Getenv("CLIENT_ADDRESS")},
    AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
    AllowHeaders: []string{},
  }))

	routes.Subscribe(e)
  routes.Altcha(e)

  e.Logger.Fatal(e.Start(":1323"))
}

