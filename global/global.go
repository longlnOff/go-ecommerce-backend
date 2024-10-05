package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/longln/go-ecommerce-backend-api/pkg/logger"
	"github.com/longln/go-ecommerce-backend-api/pkg/setting"
)


var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
)

