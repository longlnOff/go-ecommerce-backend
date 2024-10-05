package initialize


func Run() {
	// 1. Load config
	LoadConfig()
	// 2. Init logger
	InitLogger()
	// 3. Connect DB
	InitMySQL()
	// 4. Connect Redis
	InitRedis()
	// 5. Connect Router
	r := InitRouter()
	r.Run()
}