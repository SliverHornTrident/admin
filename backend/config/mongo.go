//go:build mongo

package config

import (
	"fmt"
	"strings"
)

type Mongo struct {
	Coll             string       `json:"Coll" yaml:"Coll" mapstructure:"Coll"`                                     // collection name
	Options          string       `json:"Options" yaml:"Options" mapstructure:"Options"`                            // mongodb options
	Database         string       `json:"Database" yaml:"Database" mapstructure:"Database"`                         // database name
	Username         string       `json:"Username" yaml:"Username" mapstructure:"Username"`                         // 用户名
	Password         string       `json:"Password" yaml:"Password" mapstructure:"Password"`                         // 密码
	MinPoolSize      uint64       `json:"MinPoolSize" yaml:"MinPoolSize" mapstructure:"MinPoolSize"`                // 最小连接池
	MaxPoolSize      uint64       `json:"MaxPoolSize" yaml:"MaxPoolSize" mapstructure:"MaxPoolSize"`                // 最大连接池
	SocketTimeoutMS  int64        `json:"SocketTimeoutMS" yaml:"SocketTimeoutMS" mapstructure:"SocketTimeoutMS"`    // 最大连接空闲时间
	ConnectTimeoutMS int64        `json:"ConnectTimeoutMS" yaml:"ConnectTimeoutMS" mapstructure:"ConnectTimeoutMS"` // socket超时时间
	IsZap            bool         `json:"IsZap" yaml:"IsZap" mapstructure:"IsZap"`                                  // 是否开启zap日志
	Hosts            []*MongoHost `json:"Hosts" yaml:"Hosts" mapstructure:"Hosts"`                                  // 主机列表
}

type MongoHost struct {
	Host string `json:"Host" yaml:"Host" mapstructure:"Host"` // ip地址
	Port string `json:"Port" yaml:"Port" mapstructure:"Port"` // 端口
}

// Uri .
func (x *Mongo) Uri() string {
	length := len(x.Hosts)
	hosts := make([]string, 0, length)
	for i := 0; i < length; i++ {
		if x.Hosts[i].Host != "" && x.Hosts[i].Port != "" {
			hosts = append(hosts, x.Hosts[i].Host+":"+x.Hosts[i].Port)
		}
	}
	if x.Options != "" {
		return fmt.Sprintf("mongodb://%s/%s?%s", strings.Join(hosts, ","), x.Database, x.Options)
	}
	return fmt.Sprintf("mongodb://%s/%s", strings.Join(hosts, ","), x.Database)
}
