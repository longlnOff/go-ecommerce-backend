package initialize

import (
	"fmt"
	"time"

	"github.com/longln/go-ecommerce-backend/global"
	"github.com/longln/go-ecommerce-backend/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	// init mysql
	m := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "failed to connect database")
	global.Logger.Info("connect database ok")
	global.Mdb = db

	// Set Pool
	SetPool()

	// Migrate
	migrateTables()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	checkErrorPanic(err, "set pool failed")
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns) // gioi han so luong ket noi toi da
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("migrate error: ", err)
	}
}
