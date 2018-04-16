package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sandalwing/echo-logrusmiddleware"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func main() {
	err := loadConfiguration()
	if err != nil {
		fmt.Println("An error occurred: ", err)
	} else {
		e := configEchoFramework()
		initRouter(e)

		// Server
		e.Logger.Fatal(e.Start(":" + viper.GetString("Port")))
	}
}

//initRouter will create the router for project
func initRouter(e *echo.Echo) {
	// apiRouter := e.Group("/api")
}

// configEchoFramework will return echo variable already config
func configEchoFramework() *echo.Echo {
	e := echo.New()
	e.Logger = logrusmiddleware.Logger{logrus.StandardLogger()}
	e.Use(logrusmiddleware.Hook())

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(viper.GetString("HeaderAllowOrigins"), ","),
		AllowMethods: strings.Split(viper.GetString("HeaderAllowMethods"), ","),
		AllowHeaders: strings.Split(viper.GetString("HeaderAllowHeaders"), ","),
	}))

	return e
}

// loadConfiguration will read and get config from config file, return error if something wrong
func loadConfiguration() error {
	var prefix string
	switch os.Getenv("ENVIRONMENT") {
	case "PRODUCTION":
		prefix = "prod."
	default:
		prefix = "develop."
	}

	viper.SetConfigName(prefix + "config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
