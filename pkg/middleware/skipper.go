package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	// Skipper defines a function to skip middleware. Returning true skips processing
	// the middleware.
	Skipper func(c echo.Context) bool
)

// DefaultSkipper returns false which processes the middleware.
func DefaultSkipper(c echo.Context) bool {
	if strings.Contains(c.Request().RequestURI, "/public/") ||
		strings.Contains(c.Request().RequestURI, "favicon") ||
		strings.Contains(c.Request().RequestURI, "/js/") ||
		strings.Contains(c.Request().RequestURI, "/metrics") ||
		strings.Contains(c.Request().RequestURI, "/files") ||
		strings.Contains(c.Request().Method, "OPTIONS") {
		return true
	}
	return false
}
