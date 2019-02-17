package application

import (
	"net/http"

	"mdstest/handler"

	"github.com/labstack/echo"
)

//SetupRoute performs web server routes setup
func (app *Application) SetupRoute() {

	app.Echo.GET("/", func(c echo.Context) error {
		c.Logger().Info("TEST LOGGING")
		return c.String(http.StatusOK, "It Works!")
	})

	app.Echo.POST("/User/Create", handler.UserCreate)
	app.Echo.GET("/User/:user_id", handler.UserQuery)
	app.Echo.PUT("/User/Update", handler.UserUpdate)
	app.Echo.DELETE("/User/Delete/:user_id", handler.UserDelete)

	app.Echo.POST("/UserSetting/Create", handler.UserSettingCreate)
	app.Echo.PUT("/UserSetting/Update", handler.UserSettingUpdate)
	app.Echo.DELETE("/UserSetting/Delete/:user_id/:setting_key", handler.UserSettingDelete)
}