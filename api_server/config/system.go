package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	DbType            string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	AuthCheckUrl  string `mapstructure:"auth-check-url" json:"auth-check-url" yaml:"auth-check-url"`
	FlowserverUrl string `mapstructure:"flowserver-url" json:"flowserver-url" yaml:"flowserver-url"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	IsAuthCheck   bool   `mapstructure:"is-auth-check" json:"is-auth-check" yaml:"is-auth-check"`
}
