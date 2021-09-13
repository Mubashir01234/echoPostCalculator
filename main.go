package main

import (
	"github.com/labstack/echo/v4"

	// _ "github.com/swaggo/echo-swagger/example/docs"
	_ "github.com/Mubashir01234/echoPostCalculator/docs"
	"github.com/Mubashir01234/echoPostCalculator/routes"
	_ "github.com/Mubashir01234/echoPostCalculator/routes"
)

func main() {
	e := echo.New()
	routes.Routes(e)
	e.Logger.Fatal(e.Start(":3000"))
}


