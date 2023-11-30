package main

import (
	"github.com/labstack/echo/v4"
	"numenv_subscription_api/routes"
)

func main() {
	e := echo.New()
	routes.Subscribe(e)
	e.Logger.Fatal(e.Start(":1323"))
}
