package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"mdstest/helper"
	logrusHelper "mdstest/helper/logrus"

	"github.com/BurntSushi/toml"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

//LoggerConfig represents configuration items for the logger middleware
type LoggerConfig struct {
	Skipper 	middleware.Skipper	//Skipper function (for skipping execution of middleware)
	LogRotation	RotationConfig		//config related to log rotation
	Formatter	FormatterConfig		//config related to log formatter
}

//RotationConfig represents configurations to use with logrus plugin: lestrrat-go/file-rotatelogs
type RotationConfig struct {
	LogPath				string			//path to store log file (including the log file name)
	MaxAge				Duration		//maximum age of a log file. log file older than this will be purged
	RotationInterval	Duration		//interval for rotating log files
	RotationFormat		string			//format string to use to append to log file names (e.g. "%Y-%m-%d.log")
	LinkName 			string			//string to append to the log path for the symbolic link to current log file
}

type FormatterConfig struct {
	TimeStampFormat		string			//format of time in log entry when displayed
	FullTimeStamp		bool			//config to enable/disable the entry timestamp when logging
}

//Duration is a wrapper for time.Duration for config time string to parse into time.Duration
type Duration struct {
	time.Duration
}

//UnmarshalText satisfies interface BurntSushi/toml/TextUnmarshaler
//Parses string from toml to time.Duration
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

//Logger returns a logger middleware for the echo web server logs using sirupsen/logrus library
//This is meant as a replacement for the default echo Middleware.Logger
func Logger() echo.MiddlewareFunc {

	//initiate default logger config
	defaultConfig := LoggerConfig{
		Skipper: middleware.DefaultSkipper,
	}

	//load the remaining config values from a default config file (file location is hardcoded and is relative from current working directory which should be where "main.go" is located)
	executablePath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("Failed obtaining executable path: %+v", err))
	}

	_, err = toml.DecodeFile(path.Join(executablePath, "../config/middleware/logger.toml"), &defaultConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed decoding default logger config: %+v", err))
	}

	return LoggerWithConfig(defaultConfig)
}

//LoggerWithConfig returns a logger middleware with config passed from param
func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {

	//setup logrus logger instance
	logger := logrus.New()

	//wrap logrus textFormatter with a custom formatter
	textFormatter := new(logrus.TextFormatter)
	//textFormatter.DisableColors = true
	textFormatter.DisableSorting = true

	textFormatter.FullTimestamp = config.Formatter.FullTimeStamp
	textFormatter.TimestampFormat = config.Formatter.TimeStampFormat

	customFormatter := logrusHelper.TimezoneFormatter{
		Formatter: textFormatter,
		Location: helper.DefaultLocation,
	}
	logger.Formatter = customFormatter

	logger.Out = os.Stdout //NOTE: for logging to file, the setup is done via hook, see below

	//setup hook for logrus
	logPath := config.LogRotation.LogPath
	rotateWriter, err := rotatelogs.New(
		logPath + config.LogRotation.RotationFormat,
		rotatelogs.WithLinkName(logPath + config.LogRotation.LinkName),
		rotatelogs.WithMaxAge(config.LogRotation.MaxAge.Duration),
		rotatelogs.WithRotationTime(config.LogRotation.RotationInterval.Duration),
	)

	if err != nil {
		panic(fmt.Sprintf("Failed setting up log rotation middleware: %+v", err))
	}

	logfileFormatter := logrusHelper.TimezoneFormatter{
		Formatter: &logrus.TextFormatter{
			DisableSorting: true,
			DisableColors: true,
		},
		Location: helper.DefaultLocation,
	}

	//setup logrus hook for all log levels
	logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel : rotateWriter,
			logrus.WarnLevel : rotateWriter,
			logrus.ErrorLevel: rotateWriter,
			logrus.FatalLevel: rotateWriter,
			logrus.PanicLevel: rotateWriter,
		},
		logfileFormatter,
	))

	//compose the middleware function
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			//set the echo context logger instance
			c.Echo().Logger = NewContextLogger(logger)

			start := time.Now() //mark start time for counting total processing time/latency

			//NOTE: for proper and correct measurement, this logger should be set as the first middleware used
			//middleware are basically decorators wrapping other decorators
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			latency := stop.Sub(start)

			logEntry := logger.WithFields(logrus.Fields{
				//"at"		 : time.Now().Format("2006-01-02 15:04:05"),
				"status"	 : c.Response().Status,
				"method"	 : c.Request().Method,
				"uri"		 : c.Request().RequestURI,
				"remote addr": c.Request().RemoteAddr,
				"request id" : c.Request().Header.Get(echo.HeaderXRequestID),
				"response id": c.Response().Header().Get(echo.HeaderXRequestID),
				"host"		 : c.Request().Host,
				"referer"	 : c.Request().Referer(),
				"user-agent" : c.Request().UserAgent(),
				"bytes-in"	 : c.Request().Header.Get(echo.HeaderContentLength),
				"bytes-out"	 : c.Response().Size,
				"latency"	 : latency.String(),
				"error"		 : fmt.Sprintf("%+v",err),
			})

			if c.Response().Status >= 500 {
				var errorText string
				if err != nil {
					errorText = fmt.Sprintf("%+v", err)
				}
				logEntry.Errorln(errorText)
			} else {
				//NOTE: it's possible to ignore (not logging) certain requests with http status (e.g. 404) here
				logEntry.Infoln(fmt.Sprintf("%+v %+v",c.Response().Status, c.Request().RequestURI))
			}

			//NOTE: do not pass to next middleware i.e. call next(c), as this middleware is meant to be the middleware that wraps all other middlewares
			//passing to next middleware results in duplicated output
			return nil
		}
	}
}