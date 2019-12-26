package routes

import (
	"github.com/labstack/echo"
	healthApi "github.com/trinhquocviet/go-echo-skeleton/api/health"
)

// SetupAPIRoutes is main function to start to setup routes
func SetupAPIRoutes(e *echo.Echo) {
	root := e.Group("")
	healthApi.InitRoutes(root)
}
