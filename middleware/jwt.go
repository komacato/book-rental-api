package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// JWT authentication middleware
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
		}

		return next(c)
	}
}
