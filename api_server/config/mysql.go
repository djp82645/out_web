package config

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	GormLog      int    `mapstructure:"gorm-log" json:"gorm-log" yaml:"gorm-log"`
}

