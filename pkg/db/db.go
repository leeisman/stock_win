package db

// Database ...
type Database struct {
	Debug          bool   `yaml:"debug" mapstructure:"debug"`
	Host           string `yaml:"host" mapstructure:"host"`
	Port           string `yaml:"port" mapstructure:"port"`
	Username       string `yaml:"username" mapstructure:"username"`
	Password       string `yaml:"password" mapstructure:"password"`
	DBName         string `yaml:"dbname" mapstructure:"dbname"`
	MaxIdleConns   int    `yaml:"maxidleconns" mapstructure:"maxidleconns"`
	MaxOpenConns   int    `yaml:"maxopenconns" mapstructure:"maxopenconns"`
	MaxLifetimeSec int    `yaml:"maxlifetimesec" mapstructure:"maxlifetimesec"`
	ReadTimeout    string `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout   string `yaml:"write_timeout" mapstructure:"write_timeout"`
	SearchPath     string `yaml:"search_path" mapstructure:"search_path"` // pg should setting this value. It will restrict access to db schema.
	SSLEnable      bool   `yaml:"ssl_enable" mapstructure:"ssl_enable"`   // pg ssl mode
}
