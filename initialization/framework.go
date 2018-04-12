package initialization

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sandalwing/echo-logrusmiddleware"
	"github.com/spf13/viper"
	"strings"
)

// ConfigEchoFramework will return echo variable already config
func ConfigEchoFramework() *echo.Echo {
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
