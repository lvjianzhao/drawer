package config

type License struct {
	Url           string `mapstructure:"url" json:"url" yaml:"url"`                               // 地址
	Auth          bool   `mapstructure:"auth" json:"auth" yaml:"auth"`                            // 是否开启basic认证
	Authorization string `mapstructure:"authorization" json:"authorization" yaml:"authorization"` // basic认证字符串
}
