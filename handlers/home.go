package handlers

import (
	"fmt"

	"github.com/CRSylar/hattory/view/pages"
	"github.com/labstack/echo/v4"
)

func Home(e echo.Context) error {
	// Perform any logic here
	// ...
	fmt.Printf("Pretend any server logic is handled first in Home page serving\n")
	// render the home page
	return render(e, pages.Page())
}