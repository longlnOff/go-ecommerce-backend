package initialize

import (
	"fmt"

	"github.com/longln/go-ecommerce-backend/global"
	"go.uber.org/zap"
)

func Run() {
	// Load config
	LoadConfig()
	fmt.Println("Config: ", global.Config)

	// Load logger
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))

	// Init mysql
	InitMysql()

	// Init redis
	InitRedis()

	// Init router
	r := InitRouter()
	r.Run("0.0.0.0:8080")
}