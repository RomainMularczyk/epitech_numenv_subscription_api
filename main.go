package main

import (
	"numenv_subscription_api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.Subscribe(e)

	e.Logger.Fatal(e.Start(":1323"))
}
