package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"mdstest/application"
)


func main() {
	app := &application.Application{Echo: echo.New()}

	app.SetupConfig()
	app.SetupRoute()
	app.SetupMiddleware()

	shutdownDuration, err := time.ParseDuration(app.AppConfig.Server.ShutdownTimeout)
	if err != nil {
		panic(err)
	}

	// Start server
	go func() {
		if err := app.Echo.Start(":" + strconv.Itoa(app.AppConfig.Server.ListeningPort)); err != nil {
			app.Echo.Logger.Infof("Shutting down the server: %+v", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), shutdownDuration)
	defer cancel()
	if err := app.Echo.Shutdown(ctx); err != nil {
		app.Echo.Logger.Fatal(err)
	}
}
