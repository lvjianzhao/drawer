package config

type SMTP struct {
	Username string `mapstructure:"user" json:"user" yaml:"user"`             // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Server   string `mapstructure:"server" json:"server" yaml:"server"`       // smtp服务
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	Starttls string `mapstructure:"Starttls" json:"Starttls" yaml:"Starttls"` // 是否开启tls
}
