package health

import (
	"github.com/labstack/echo"
)

// InitRoutes config for Controller
func InitRoutes(group *echo.Group) {
	apiRoutes := group.Group("/auth")
	var ctr Controller

	apiRoutes.GET("/health-check", ctr.Check)
}
