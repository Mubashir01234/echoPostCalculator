package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	// _ "github.com/swaggo/echo-swagger/example/docs"
	_ "github.com/Mubashir01234/echoPostCalculator/docs"
)

type Numbers struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
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
	return c.String(http.StatusOK, "Number1 + Number2 = "+strconv.Itoa(add))
}
func subtraction(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sub := n.Number1 - n.Number2
	return c.String(http.StatusOK, "Number1 - Number2 = "+strconv.Itoa(sub))
}
func multiplication(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mul := n.Number1 * n.Number2
	return c.String(http.StatusOK, "Number1 * Number2 = "+strconv.Itoa(mul))
}
func division(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	div := float64(n.Number1) / float64(n.Number2)
	return c.String(http.StatusOK, "Number1 / Number2 = "+strconv.FormatFloat(div, 'E', -1, 64))
}
func modules(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mod := n.Number1 % n.Number2
	return c.String(http.StatusOK, "Number1 % Number2 = "+strconv.Itoa(mod))
}
func power(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	power := math.Pow(float64(n.Number1), float64(n.Number2))
	pow := int(power)
	return c.String(http.StatusOK, "Number1 ^ Number2 = "+strconv.Itoa(pow))
}
func square(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	squ1 := n.Number1 * n.Number1
	squ2 := n.Number2 * n.Number2
	return c.String(http.StatusOK, "Square of Number1 is: "+strconv.Itoa(squ1)+"\n"+"Square of Number2 is: "+strconv.Itoa(squ2))
}
func squareRoot(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sqrt1 := math.Sqrt(float64(n.Number1))
	sqrt2 := math.Sqrt(float64(n.Number2))

	return c.String(http.StatusOK, "Square Root of Number1 = "+strconv.FormatFloat(sqrt1, 'E', -1, 64)+"\nSquare Root of Number2 = "+strconv.FormatFloat(sqrt2, 'E', -1, 64))
}
