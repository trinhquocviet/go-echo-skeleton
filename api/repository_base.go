package api

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	constants "github.com/trinhquocviet/go-echo-skeleton/constants"
)

// RepositoryBase the base class for api API class
type RepositoryBase struct {
	DB *gorm.DB
}

// OpenConnection the function to init the db connection and assign to Base.DB
func (inst *RepositoryBase) OpenConnection() (err error) {
	cnStr := fmt.Sprintf(
		`%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		constants.Env["DB_USER"],
		constants.Env["DB_PASS"],
		constants.Env["DB_HOST"],
		constants.Env["DB_PORT"],
		constants.Env["DB_NAME"],
	)

	inst.DB, err = gorm.Open("mysql", cnStr)
	return
}