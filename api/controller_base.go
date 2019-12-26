package api

import (
	"github.com/labstack/echo"
	msg "github.com/trinhquocviet/go-echo-skeleton/internal/message"
	"net/http"
	"reflect"
)

type (
	// DefaultSuccessRes the default response
	DefaultSuccessRes struct {
		Success bool `json:"Success"`
	}

	// DefaultErrorRes the default response
	DefaultErrorRes struct {
		Success bool   `json:"Success"`
		Error   string `json:"Error"`
	}

	// ControllerBase the base class for api API class
	ControllerBase struct {
		repository RepositoryBase
	}
)


func (inst *ControllerBase) getHeaderValue(c echo.Context, key string) string {
	return c.Request().Header.Get(key)
}

// parse will parse args to code and response data
func (inst *ControllerBase) parse(args []interface{}) (code int, res interface{}) {
	if len(args) > 0 {
		count := len(args)
		if count == 1 {
			args0 := reflect.TypeOf(args[0]).String()
			if args0 == "int" {
				code = args[0].(int)
			} else if args0 == "string" || args0 == "errors.errorString" {
				res = DefaultErrorRes{Success: false, Error: args[0].(string)}
			} else if reflect.TypeOf(args[0]) != nil {
				res = args[0]
			}
		} else if count == 2 {
			args0 := reflect.TypeOf(args[0]).String()
			args1 := reflect.TypeOf(args[1]).String()
			if args0 == "int" {
				code = args[0].(int)
			}
			if args1 == "string" || args1 == "*errors.errorString" {
				res = DefaultErrorRes{Success: false, Error: args[0].(string)}
			} else if reflect.TypeOf(args[1]) != nil {
				res = args[1]
			}
		}
	}

	return
}

// SuccessRes return the json response
func (inst *ControllerBase) SuccessRes(c echo.Context, args ...interface{}) error {
	code, res := inst.parse(args)

	if code < 200 {
		code = http.StatusOK
	}
	if res == nil {
		res = DefaultSuccessRes{
			Success: true,
		}
	}

	return c.JSON(code, res)
}

// ErrorRes return the json response
func (inst *ControllerBase) ErrorRes(c echo.Context, args ...interface{}) error {
	code, res := inst.parse(args)

	if code < 200 {
		code = http.StatusBadRequest
	}
	if res == nil {
		res = DefaultErrorRes{
			Success: false,
			Error:   msg.DefaultError,
		}
	}

	return c.JSON(code, res)
}
