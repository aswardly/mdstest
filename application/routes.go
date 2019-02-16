package application

import (
	"net/http"

	"github.com/labstack/echo"
)

//SetupRoute performs web server routes setup
func (app *Application) SetupRoute() {

	app.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It Works!")
	})
}