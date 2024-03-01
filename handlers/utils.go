package handlers

import (
	"github.com/CRSylar/hattory/view/pages"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(e echo.Context, c templ.Component) error {
	// render the home page
	layoutWrap := pages.Layout(c)
	return layoutWrap.Render(e.Request().Context(), e.Response())
}