package global

import (
	"database/sql"
	"github.com/longln/go-ecommerce-backend-api/pkg/logger"
	"github.com/longln/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
)


var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
	Mdbc *sql.DB
)

