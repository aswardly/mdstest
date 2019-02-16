package application

import (
	"mdstest/middleware"
)

//SetupMiddleware performs web server middleware setup
func (app *Application) SetupMiddleware() {
	app.Echo.Use(middleware.Logger()) //NOTE: must be the first middleware set
	app.Echo.Use(middleware.Registry())
	app.Echo.Use(middleware.Recover())
}