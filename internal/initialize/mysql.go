package initialize

import (
	"fmt"
	"time"

	"github.com/longln/go-ecommerce-backend/global"
	"github.com/longln/go-ecommerce-backend/internal/model"
	// "github.com/longln/go-ecommerce-backend/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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

	genTableDAO()
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
		// &po.User{},
		// &po.Role{},
		&model.UserRole2{},
	)
	if err != nil {
		fmt.Println("migrate error: ", err)
	}
}

func genTableDAO() {
	// Initiate the tables
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	  })
	
	  // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	// g.GenerateAllTable()
	g.GenerateModel("user_roles")

	//   // Generate basic type-safe DAO API for struct `model.User` following conventions
	//   g.ApplyBasic(model.User{})
	
	//   // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//   g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})
	
	  // Generate the code
	  g.Execute()
}
