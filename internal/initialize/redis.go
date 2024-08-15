package initialize

import (
	"context"
	"fmt"

	"github.com/longln/go-ecommerce-backend/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)


var ctx = context.Background()
func InitRedis() {
	r := global.Config.Redis
	// init redis
	rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
    })

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("redis connect error", zap.Error(err))
	}
	fmt.Println("redis connect ok")
	global.Rdb = rdb
}