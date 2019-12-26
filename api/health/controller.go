package health

import (
	"github.com/labstack/echo"
	api "github.com/trinhquocviet/go-echo-skeleton/api"
	internal "github.com/trinhquocviet/go-echo-skeleton/internal"
)

// Controller the handler function extend from api/controller_base.go
type Controller struct {
	api.ControllerBase
	repository Repository
	validator  internal.Validator
}

// Check health-check
func (inst *Controller) Check(c echo.Context) error {
	// Res
	var res CheckRes
	res.Status = "ok"

	return inst.SuccessRes(c, res)
}