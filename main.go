package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	// _ "github.com/swaggo/echo-swagger/example/docs"
	_ "github.com/Mubashir01234/echoPostCalculator/docs"
)

type Numbers struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type Response struct {
	Result float64 `json:"result"`
}

func main() {
	e := echo.New()
	e.POST("/calculator/add", addition)
	e.POST("/calculator/sub", subtraction)
	e.POST("/calculator/mul", multiplication)
	e.POST("/calculator/div", division)
	e.POST("/calculator/mod", modules)
	e.POST("/calculator/power", power)
	e.POST("/calculator/square", square)
	e.POST("/calculator/squareroot", squareRoot)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
func addition(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	add := n.Number1 + n.Number2
	result:=Response{
		add,
	}
	return c.JSON(http.StatusOK,result)
	// return c.String(http.StatusOK, "Number1 + Number2 = "+strconv.Itoa(add))
}
func subtraction(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sub := n.Number1 - n.Number2
	result:=Response{
		sub,
	}
	return c.JSON(http.StatusOK, result)
}
func multiplication(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mul := n.Number1 * n.Number2
	result:=Response{
		mul,
	}
	return c.JSON(http.StatusOK, result)
}
func division(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	div := n.Number1 / n.Number2
	result:=Response{
		div,
	}
	return c.JSON(http.StatusOK, result)
}
func modules(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mod := int(n.Number1) % int(n.Number2)
	result:=Response{
		float64(mod),
	}
	return c.JSON(http.StatusOK, result)
}
func power(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	power := math.Pow(n.Number1,n.Number2)
	result:=Response{
		power,
	}
	return c.JSON(http.StatusOK, result)
}
func square(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	squ := n.Number1 * n.Number1
	result:=Response{
		squ,
	}
	return c.JSON(http.StatusOK, result)
}
func squareRoot(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sqrt := math.Sqrt(n.Number1)
	result:=Response{
		sqrt,
	}
	return c.JSON(http.StatusOK, result)
}
