package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Sever struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
		Port int `mapstructure:"port"`
		Name string `mapstructure:"name"`
	}

	SecurityJWT struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("config/")	// path to config
	viper.SetConfigName("local")			// file name
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file %w", err))
	}

	// read port
	fmt.Println("Server Port: ", viper.GetInt("server.port"))
	fmt.Println("Server Security: ", viper.GetString("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to read configuration file %w", err))
	}

	fmt.Println(config)
	fmt.Println("Config port: ", config.Sever.Port)
	fmt.Println("Config Security: ", config.SecurityJWT.JWT.Key)

	for _, db := range config.Databases {
		fmt.Printf("databases User: %s, Password: %s, Host: %s, Port: %d, Name: %s\n", db.User, db.Password, db.Host, db.Port, db.Name)
	}
}