package initialize

import (
	"github.com/longln/go-ecommerce-backend/global"
	"github.com/longln/go-ecommerce-backend/pkg/logger"
)

func InitLogger() {
	// init logger
	global.Logger = logger.NewLogger(global.Config.Logger)
}