package routes

import (
	//_ "github.com/Mubashir01234/echoPostCalculator/controller"
	"github.com/Mubashir01234/echoPostCalculator/controller"
	middlewares "github.com/Mubashir01234/echoPostCalculator/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo) {

	// e.Use(middlewares.MubashirMiddleware)
	e.POST("/calculator/add", controller.Addition, middlewares.MubashirMiddleware)
	e.POST("/calculator/sub", controller.Subtraction, middlewares.MubashirMiddleware)
	e.POST("/calculator/mul", controller.Multiplication, middlewares.MubashirMiddleware)
	e.POST("/calculator/div", controller.Division, middlewares.MubashirMiddleware)
	e.POST("/calculator/mod", controller.Modulus, middlewares.MubashirMiddleware)
	e.POST("/calculator/power", controller.Power, middlewares.MubashirMiddleware)
	e.POST("/calculator/square", controller.Square, middlewares.MubashirMiddleware)
	e.POST("/calculator/squareroot", controller.SquareRoot, middlewares.MubashirMiddleware)
	e.GET("/calculator/getRecord/:id", controller.GetRecord, middlewares.MubashirMiddleware)
	e.GET("/calculator/getAllRecord", controller.GetAllRecord, middlewares.MubashirMiddleware)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
