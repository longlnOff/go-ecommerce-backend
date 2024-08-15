package setting

type Config struct {
	Sever ServerSetting `mapstructure:"server"`
	MySQL MySQLSetting `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis RedisSetting `mapstructure:"redis"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	LogLevel string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxSize int `mapstructure:"maxSize"`
	MaxBackups int `mapstructure:"maxBackups"`
	MaxAge int `mapstructure:"maxAge"`
	Compress bool `mapstructure:"compress"`
}

type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int `mapstructure:"database"`
}