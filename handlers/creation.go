package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func CreateRequestsHandler() http.Handler {

	// ...
	e = echo.New()

	e.GET("/lib/htmx.js", func(c echo.Context) error {
		if strings.Contains(c.Request().Header.Get("Accept-Encoding"), "gzip") {
			c.Response().Header().Set("Content-Encoding", "gzip")
			c.Response().Header().Set("Content-Type", "application/gzip")
			return c.File("lib/htmx/htmx.min.js.gz")
		}
		c.Response().Header().Set("Content-Type", "application/javascript")
		return c.File("lib/htmx/htmx.min.js")
	})

	e.GET("/lib/_hyperscript.js", func(c echo.Context) error {
		if strings.Contains(c.Request().Header.Get("Accept-Encoding"), "gzip") {
			c.Response().Header().Set("Content-Encoding", "gzip")
			c.Response().Header().Set("Content-Type", "application/gzip")
			return c.File("lib/_hyperscript/_hyperscript.min.js.gz")
		}
		c.Response().Header().Set("Content-Type", "application/javascript")
		return c.File("lib/_hyperscript/_hyperscript.min.js")
	})

	e.GET("/view/public/styles.css", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/css")
		return c.File("view/public/styles.css")
	})

	return e
}
