package main

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )



// func main() {
// 	viper := viper.New()

// 	// set config path
// 	viper.AddConfigPath("config")
// 	// set config file name
// 	viper.SetConfigName("local")
// 	// set config file type
// 	viper.SetConfigType("yaml")

// 	// Find and read config file (must do)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic("Can't find config file")
// 	}
	
// 	// read config
// 	if err := viper.Unmarshal(&config); err != nil {
// 		fmt.Printf("Unable to decode configuration %v", err)
// 	}

// 	// read server configuration
// 	fmt.Println("Config Port::", config.Server.Port)

// 	for _, db := range(config.Databases) {
// 		fmt.Printf("databases User: %s, password: %s, host: %s\n", db.User, db.Password, db.Host)
// 	}
// }