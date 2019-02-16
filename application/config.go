package application

import (
	"mdstest/application/config"

	"fmt"
	"os"
	"path"

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
}
