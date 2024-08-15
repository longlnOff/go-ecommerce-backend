package initialize

import (
	"fmt"

	"github.com/longln/go-ecommerce-backend/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// load config
	viper := viper.New()
	viper.AddConfigPath("config/")	// path to config
	viper.SetConfigName("local")			// file name
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to read configuration file %w", err))
	}
}