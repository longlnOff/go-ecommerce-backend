package global

import (
	"github.com/longln/go-ecommerce-backend/pkg/logger"
	"github.com/longln/go-ecommerce-backend/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
	Rdb *redis.Client
)