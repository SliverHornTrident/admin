//go:build elasticsearch

package config

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"net"
	"net/http"
	"time"
)

type Elasticsearch struct {
	Username string   `json:"Username" yaml:"Username" mapstructure:"Username"` // 用户名
	Password string   `json:"Password" yaml:"Password" mapstructure:"Password"` // 密码
	Address  []string `json:"Address" yaml:"Address" mapstructure:"Address"`    // 地址
}

func (c *Elasticsearch) Config() elasticsearch.Config {
	return elasticsearch.Config{
		Addresses: c.Address,
		Username:  c.Username,
		Password:  c.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig:       &tls.Config{MinVersion: tls.VersionTLS12},
		},
	}
}
