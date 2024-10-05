package setting


type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Databases []struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host	string `mapstructure:"host"`
	} `mapstructure:"databases"`
	MySQL MySQLSetting `mapstructure:"mysql"`
	Logger LogSetting `mapstructure:"logger"`
	Redis RedisSetting `mapstructure:"redis"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}



type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}


type MySQLSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`
}

type LogSetting struct {
	LogLevel string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize int `mapstructure:"MaxSize"`
	MaxBackups int `mapstructure:"MaxBackups"`
	MaxAge int `mapstructure:"MaxAge"`
	Compress bool `mapstructure:"Compress"`
}