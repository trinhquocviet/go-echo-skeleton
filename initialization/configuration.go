package initialization

import (
	"github.com/spf13/viper"
	"os"
)

// LoadConfiguration will read and get config from config file, return error if something wrong
func LoadConfiguration() error {
	var prefix string
	switch os.Getenv("ENVIRONMENT") {
	case "PRODUCTION":
		prefix = "prod."
	default:
		prefix = "develop."
	}

	viper.SetConfigName(prefix + "config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
