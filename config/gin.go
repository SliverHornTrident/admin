//go:build gin

package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Port     int    `json:"Port" yaml:"Port" mapstructure:"Port"`
	Node     string `json:"Node" yaml:"Node" mapstructure:"Node"`
	Host     string `json:"Host" yaml:"Host" mapstructure:"Host"`
	Name     string `json:"Name" yaml:"Name" mapstructure:"Name"`
	Graceful bool   `json:"Graceful" yaml:"Graceful" mapstructure:"Graceful"`
}

func (c *Gin) Address() string {
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.Host == "" {
		return fmt.Sprintf(":%d", c.Port)
	}
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Gin) Etcd() string {
	return "/" + c.Name
}

func (c *Gin) EtcdBinary() string {
	return fmt.Sprintf("%s_%s", c.Name, gin.Mode())
}

func (c *Gin) EtcdCommand() string {
	if c.Node != "" {
		return fmt.Sprintf("systemctl restart %s.%s.service", c.Name, c.Node)
	}
	return fmt.Sprintf("systemctl restart %s.service", c.Name)
}
