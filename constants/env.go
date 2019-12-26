package constants

import "os"

// Env list all variable read from os enviroment
var Env = map[string]string{
	"DB_USER":              os.Getenv("DB_USER"),
	"DB_PASS":              os.Getenv("DB_PASS"),
	"DB_HOST":              os.Getenv("DB_HOST"),
	"DB_PORT":              os.Getenv("DB_PORT"),
	"DB_NAME":              os.Getenv("DB_NAME"),
	"PORT":                 os.Getenv("API_PORT"),
	"HEADER_ALLOW_ORIGINS": os.Getenv("API_HEADER_ALLOW_ORIGINS"),
	"HEADER_ALLOW_METHODS": os.Getenv("API_HEADER_ALLOW_METHODS"),
	"HEADER_ALLOW_HEADERS": os.Getenv("API_HEADER_ALLOW_HEADERS"),
}
