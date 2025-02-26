package config

import "time"

type Etcd struct {
	Username    string        `json:"Username" yaml:"Username" mapstructure:"Username"`          // 用户名
	Password    string        `json:"Password" yaml:"Password" mapstructure:"Password"`          // 密码
	DialTimeout time.Duration `json:"DialTimeout" yaml:"DialTimeout" mapstructure:"DialTimeout"` // 链接超时时间
	Endpoints   []string      `json:"Endpoints" yaml:"Endpoints" mapstructure:"Endpoints"`       // 地址
}
