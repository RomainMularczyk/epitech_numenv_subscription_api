package main

import (
	"numenv_subscription_api/routes"
	"numenv_subscription_api/services/discord"

	"github.com/labstack/echo/v4"
)

func main() {
  go func() {
    discord.DiscordClient()
  }()

  // Starting server
	e := echo.New()
	routes.Subscribe(e)
  
  e.Logger.Fatal(e.Start(":1323"))
}

