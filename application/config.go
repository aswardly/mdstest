package application

import (
	"fmt"
	"os"
	"path"

	"mdstest/application/config"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
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

	//set default echo logger
	appLogger := logrus.New()
	appLogger.SetOutput(os.Stdout)
	appLogger.SetLevel(logrus.DebugLevel)

	app.Echo.Logger = NewLogger(appLogger)
}
