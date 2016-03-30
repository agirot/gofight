package example

import (
	"github.com/labstack/echo"
	"net/http"
)

// Handler
func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func EchoEngine() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Routes
	e.Get("/", hello())

	return e
}