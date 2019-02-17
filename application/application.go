package application

import (
	"github.com/labstack/echo"
	"mdstest/application/config"
)

//Application struct represents the web application
type Application struct {
	Echo 			*echo.Echo  			//echo app instance
	AppConfig		*config.AppConfig		//application config
}