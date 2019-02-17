package application

import (
	"fmt"
	"github.com/labstack/echo"
	"mdstest/response"
	"net/http"
	"os"
	"path"

	"mdstest/application/config"

	"github.com/BurntSushi/toml"
)

//setupConfig sets application configurations
func (app *Application) SetupConfig() {
	app.AppConfig = new(config.AppConfig)

	executablePath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("Failed obtaining executable path: %+v", err))
	}

	_, err = toml.DecodeFile(path.Join(executablePath, "../config/app.toml"), app.AppConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed decoding default application config: %+v", err))
	}

	//set custom echo error handler
	app.Echo.HTTPErrorHandler = customHTTPErrorHandler
}

//customHTTPErrorHandler handles and write output in format
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	output := response.GenericResponse{
		ResponseCode: response.ResponseCodeError,
		ResponseMessage: err.Error(),
	}
	if err := c.JSON(code, output); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
