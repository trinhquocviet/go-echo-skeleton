package api

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	defaultSuccessResJSON = `{"Success":true}`
	defaultErrorResJSON   = `{"Success":false,"Error":"an error has occurred. please try again later"}`
)

func TestControllerBase_SuccessRes(t *testing.T) {
	customRes := struct {
		Transactions []string `json:"transactions"`
	}{
		Transactions: []string{"tx0001", "tx0002"},
	}
	wantResJSONForCustomRes := `{"transactions":["tx0001","tx0002"]}`

	tests := []struct {
		name             string
		customHTTPStatus int
		customRes        interface{}
		wantErr          bool
		wantHTTPStatus   int
		wantResJSON      string
	}{
		{
			name:             "Default success",
			wantErr:          false,
			customHTTPStatus: 0,
			customRes:        nil,
			wantHTTPStatus:   http.StatusOK,
			wantResJSON:      defaultSuccessResJSON,
		},
		{
			name:             "With custom HTTPStatus",
			wantErr:          false,
			customHTTPStatus: http.StatusAccepted,
			customRes:        nil,
			wantHTTPStatus:   http.StatusAccepted,
			wantResJSON:      defaultSuccessResJSON,
		},
		{
			name:             "With custom HTTPStatus and custom Response",
			wantErr:          false,
			customHTTPStatus: http.StatusAccepted,
			customRes:        customRes,
			wantHTTPStatus:   http.StatusAccepted,
			wantResJSON:      wantResJSONForCustomRes,
		},
		{
			name:             "With custom response",
			wantErr:          false,
			customHTTPStatus: 0,
			customRes:        customRes,
			wantHTTPStatus:   http.StatusOK,
			wantResJSON:      wantResJSONForCustomRes,
		},
		{
			name:             "With custom HTTPStatus and custom Response",
			wantErr:          false,
			customHTTPStatus: http.StatusOK,
			customRes:        customRes,
			wantHTTPStatus:   http.StatusOK,
			wantResJSON:      wantResJSONForCustomRes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := ControllerBase{}

			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{}`))
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)

			// Assertions
			if tt.customHTTPStatus == 0 && tt.customRes == nil {
				if assert.NoError(t, ctr.SuccessRes(c)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			} else if tt.customHTTPStatus > 0 && tt.customRes == nil {
				if assert.NoError(t, ctr.SuccessRes(c, tt.customHTTPStatus)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			} else if tt.customHTTPStatus == 0 && tt.customRes != nil {
				if assert.NoError(t, ctr.SuccessRes(c, tt.customRes)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			} else if tt.customHTTPStatus > 0 && tt.customRes != nil {
				if assert.NoError(t, ctr.SuccessRes(c, tt.customHTTPStatus, tt.customRes)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			}
		})
	}
}

func TestControllerBase_ErrorRes(t *testing.T) {
	customRes := struct {
		Success bool   `json:"Success"`
		Error   string `json:"Error"`
	}{
		Success: false,
		Error:   "custom response error",
	}
	wantResJSONForCustomRes := `{"Success":false,"Error":"custom response error"}`

	tests := []struct {
		name             string
		customHTTPStatus int
		customRes        interface{}
		wantErr          bool
		wantHTTPStatus   int
		wantResJSON      string
	}{
		{
			name:             "Default error",
			wantErr:          false,
			customHTTPStatus: 0,
			customRes:        nil,
			wantHTTPStatus:   http.StatusBadRequest,
			wantResJSON:      defaultErrorResJSON,
		},
		{
			name:             "With custom HTTPStatus",
			wantErr:          false,
			customHTTPStatus: http.StatusNotAcceptable,
			customRes:        nil,
			wantHTTPStatus:   http.StatusNotAcceptable,
			wantResJSON:      defaultErrorResJSON,
		},
		{
			name:             "With custom HTTPStatus and custom Response",
			wantErr:          false,
			customHTTPStatus: http.StatusNotAcceptable,
			customRes:        customRes,
			wantHTTPStatus:   http.StatusNotAcceptable,
			wantResJSON:      wantResJSONForCustomRes,
		},
		{
			name:             "With custom response",
			wantErr:          false,
			customHTTPStatus: 0,
			customRes:        customRes,
			wantHTTPStatus:   http.StatusBadRequest,
			wantResJSON:      wantResJSONForCustomRes,
		},
		{
			name:             "With custom HTTPStatus and custom Response",
			wantErr:          false,
			customHTTPStatus: http.StatusNotAcceptable,
			customRes:        customRes,
			wantHTTPStatus:   http.StatusNotAcceptable,
			wantResJSON:      wantResJSONForCustomRes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := ControllerBase{}

			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{}`))
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)

			// Assertions
			if tt.customHTTPStatus == 0 && tt.customRes == nil {
				if assert.NoError(t, ctr.ErrorRes(c)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			} else if tt.customHTTPStatus > 0 && tt.customRes == nil {
				if assert.NoError(t, ctr.ErrorRes(c, tt.customHTTPStatus)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			} else if tt.customHTTPStatus == 0 && tt.customRes != nil {
				if assert.NoError(t, ctr.ErrorRes(c, tt.customRes)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			} else if tt.customHTTPStatus > 0 && tt.customRes != nil {
				if assert.NoError(t, ctr.ErrorRes(c, tt.customHTTPStatus, tt.customRes)) {
					assert.Equal(t, tt.wantHTTPStatus, rec.Code)
					assert.Equal(t, tt.wantResJSON, rec.Body.String())
				}
			}
		})
	}
}

func TestControllerBase_parse(t *testing.T) {
	funcName := "ControllerBase.parse()"

	tests := []struct {
		name     string
		args     []interface{}
		wantCode int
		wantRes  interface{}
	}{
		{
			name:     "No input",
			args:     []interface{}{},
			wantCode: 0,
			wantRes:  nil,
		},
		{
			name:     "Custom HTTP status",
			args:     []interface{}{http.StatusOK},
			wantCode: http.StatusOK,
			wantRes:  nil,
		},
		{
			name:     "Custom response",
			args:     []interface{}{DefaultSuccessRes{Success: true}},
			wantCode: 0,
			wantRes:  DefaultSuccessRes{Success: true},
		},
		{
			name:     "Custom HTTP status and response",
			args:     []interface{}{http.StatusAccepted, DefaultSuccessRes{Success: true}},
			wantCode: http.StatusAccepted,
			wantRes:  DefaultSuccessRes{Success: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := ControllerBase{}
			gotCode, gotRes := ctr.parse(tt.args)

			if gotCode != tt.wantCode {
				t.Errorf("%s gotCode = %v, want %v", funcName, gotCode, tt.wantCode)
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("%s gotRes = %v, want %v", funcName, gotRes, tt.wantRes)
			}
		})
	}
}
