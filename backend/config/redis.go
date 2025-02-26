//go:build redis

package config

type Redis struct {
	Db        int      `json:"Db" yaml:"Db" mapstructure:"Db"`                      // 数据库
	Address   string   `json:"Address" yaml:"Address" mapstructure:"Address"`       // 地址
	Username  string   `json:"Username" yaml:"Username" mapstructure:"Username"`    // 用户名
	Password  string   `json:"Password" yaml:"Password" mapstructure:"Password"`    // 密码
	Addresses []string `json:"Addresses" yaml:"Addresses" mapstructure:"Addresses"` // 集群地址
}
