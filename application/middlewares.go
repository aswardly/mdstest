package application

import (
	"mdstest/middleware"
)

//SetupMiddleware performs web server middleware setup
func (app *Application) SetupMiddleware() {
	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())
}