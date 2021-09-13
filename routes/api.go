package routes

import (
	"github.com/Mubashir01234/echoPostCalculator/controller"
	_ "github.com/Mubashir01234/echoPostCalculator/controller"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo){
	e.POST("/calculator/add", controller.Addition)
	e.POST("/calculator/sub", controller.Subtraction)
	e.POST("/calculator/mul", controller.Multiplication)
	e.POST("/calculator/div", controller.Division)
	e.POST("/calculator/mod", controller.Modulus)
	e.POST("/calculator/power", controller.Power)
	e.POST("/calculator/square", controller.Square)
	e.POST("/calculator/squareroot", controller.SquareRoot)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}