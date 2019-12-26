package internal

import (
	"fmt"
	logrus "github.com/Sirupsen/logrus"
	msg "github.com/trinhquocviet/go-echo-skeleton/internal/message"
	v "gopkg.in/go-playground/validator.v9"
)

// Validator struct with custom validate
type Validator struct {
	validator *v.Validate
}

// Validate function
func (inst *Validator) Validate(i interface{}) error {
	inst.validator = v.New()
	err := inst.validator.Struct(i)
	if err != nil {
		logrus.Error("Validate Error: ", err)
		return fmt.Errorf(msg.RequestParamsMismatch)
	}
	return err
}
