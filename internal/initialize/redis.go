package initialize

import (
	"context"
	"fmt"

	"github.com/longln/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)


var ctx = context.Background()
func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Db,
		PoolSize: r.PoolSize,
    })

	_, err := rdb.Ping(ctx).Result()


	if err != nil {
		global.Logger.Error("Redis intialization Error", zap.Error(err))
	}
	fmt.Println("init redis is running")
	global.Rdb = rdb

	redisExample()

}


func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	val, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	global.Logger.Info("value score is: ", zap.String("score", val))
}