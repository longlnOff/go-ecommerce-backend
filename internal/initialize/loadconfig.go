package initialize

import (
	"fmt"

	"github.com/longln/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()

	// set config path
	viper.AddConfigPath("config")
	// set config file name
	viper.SetConfigName("local")
	// set config file type
	viper.SetConfigType("yaml")

	// Find and read config file (must do)
	err := viper.ReadInConfig()
	if err != nil {
		panic("Can't find config file")
	}
	
	// read config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}