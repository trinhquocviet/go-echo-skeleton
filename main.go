package main

import (
  "fmt"
  "github.com/Sirupsen/logrus"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/sandalwing/echo-logrusmiddleware"
  "github.com/trinhquocviet/go-echo-skeleton/routes"
  "strings"
  constants "github.com/trinhquocviet/go-echo-skeleton/constants"
)

func main() {
  e := configEchoFramework()
  routes.SetupAPIRoutes(e)

  // Server
  serverPort := fmt.Sprintf(":%s", constants.Env["PORT"])
  e.Logger.Fatal(e.Start(serverPort))
}

func configEchoFramework() *echo.Echo {
  e := echo.New()
  e.Logger = logrusmiddleware.Logger{logrus.StandardLogger()}
  e.Use(logrusmiddleware.Hook())

  // Middleware
  e.Use(middleware.Recover())
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: strings.Split(constants.Env["HEADER_ALLOW_ORIGINS"], ","),
    AllowMethods: strings.Split(constants.Env["HEADER_ALLOW_METHODS"], ","),
    AllowHeaders: strings.Split(constants.Env["HEADER_ALLOW_HEADERS"], ","),
  }))

  return e
}
