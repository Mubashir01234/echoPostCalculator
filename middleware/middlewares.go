package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ServerHeader middleware adds a `Server` header to the response.
func MubashirMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	// fmt.Println("MUBASHIRs")
	return func(c echo.Context) error {
		if (len(c.Request().Header["Authorization"]) > 0) {
			if (c.Request().Header["Authorization"][0] == "secretkey") {
				c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
				return next(c)
			}
		}
		return c.JSON(http.StatusForbidden, "You are not authorized!")
	}
}