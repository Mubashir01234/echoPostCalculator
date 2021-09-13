package controller

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)
type Numbers struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type Response struct {
	Result float64 `json:"result"`
}
func Addition(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	add := n.Number1 + n.Number2
	result := Response{
		add,
	}
	return c.JSON(http.StatusOK, result)
	// return c.String(http.StatusOK, "Number1 + Number2 = "+strconv.Itoa(add))
}
func Subtraction(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sub := n.Number1 - n.Number2
	result := Response{
		sub,
	}
	return c.JSON(http.StatusOK, result)
}
func Multiplication(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mul := n.Number1 * n.Number2
	result := Response{
		mul,
	}
	return c.JSON(http.StatusOK, result)
}
func Division(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	div := n.Number1 / n.Number2
	result := Response{
		div,
	}
	return c.JSON(http.StatusOK, result)
}
func Modulus(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mod := int(n.Number1) % int(n.Number2)
	result := Response{
		float64(mod),
	}
	return c.JSON(http.StatusOK, result)
}
func Power(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	power := math.Pow(n.Number1, n.Number2)
	result := Response{
		power,
	}
	return c.JSON(http.StatusOK, result)
}
func Square(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	squ := n.Number1 * n.Number1
	result := Response{
		squ,
	}
	return c.JSON(http.StatusOK, result)
}
func SquareRoot(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sqrt := math.Sqrt(n.Number1)
	result := Response{
		sqrt,
	}
	return c.JSON(http.StatusOK, result)
}