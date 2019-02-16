package middleware

import (
	"fmt"
	"os"
	"path"

	"mdstest/application/config"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RegistryConfig defines config for the Registry middleware
type RegistryConfig struct {
	Skipper 	middleware.Skipper	//Skipper function (for skipping execution of middleware)
	Config 		*config.AppConfig	//Application config
}

// Registry returns a middleware which sets dependencies in the echo context (as a registry) using default config
func Registry() echo.MiddlewareFunc {

	defaultAppConfig := new(config.AppConfig)

	//parse application config
	executablePath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("Failed obtaining executable path: %+v", err))
	}

	_, err = toml.DecodeFile(path.Join(executablePath, "../config/app.toml"), defaultAppConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed decoding default application config: %+v", err))
	}

	registryConfig := RegistryConfig{Config: defaultAppConfig}
	return RegistryWithConfig(registryConfig)
}

// RegistryWithConfig returns a middleware which sets dependencies in the echo context (as a registry) with injected config
func RegistryWithConfig(config RegistryConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultRecoverConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			//setup gorm
			db, err := gorm.Open(config.Config.Database.DriverType, config.Config.Database.DSN)
			if err != nil {
				c.Logger().Fatalf("Failed creating Gorm instance: %+v\n", err)
				return err
			}
			c.Set("GORM",db)

			//TODO: setup other required objects here

			return next(c)
		}
	}
}