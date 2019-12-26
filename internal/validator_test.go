package internal

import (
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	funcName := "Validate()"

	tests := []struct {
		Name    string
		Obj     interface{}
		WantErr bool
	}{
		{
			Name: "Validate empty required",
			Obj: struct {
				code int `validate:"required,max=200"`
			}{},
			WantErr: true,
		},
		{
			Name: "Validate code more than 200",
			Obj: struct {
				code int `validate:"required,max=200"`
			}{code: 201},
			WantErr: true,
		},
		{
			Name: "Validate without validate config",
			Obj: struct {
				code int
			}{code: 201},
			WantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			inst := &Validator{}
			if err := inst.Validate(tt.Obj); err != nil && tt.WantErr {
				t.Errorf("%v error = %v, wantErr %v", funcName, err, tt.WantErr)
			}
		})
	}
}
