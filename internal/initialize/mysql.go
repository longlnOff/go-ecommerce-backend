package initialize

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/longln/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)





func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
	)
	db, err := sql.Open("mysql", dsn)
	checkErrorPanic(err, "MySQL intialization Error")
	global.Logger.Info("MySQL intialization success")
	global.Mdbc = db

	SetPool()
}


func SetPool() {
	m := global.Config.MySQL
	global.Mdbc.SetMaxOpenConns(m.MaxOpenConns)
	global.Mdbc.SetMaxIdleConns(m.MaxIdleConns)
	global.Mdbc.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}