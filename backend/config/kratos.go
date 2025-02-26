//go:build kratos

package config

import (
	"fmt"
	"time"
)

type Kratos struct {
	Id      string     `json:"Id" yaml:"Id" mapstructure:"Id"`
	Name    string     `json:"Name" yaml:"Name" mapstructure:"Name"`
	Version string     `json:"Version" yaml:"Version" mapstructure:"Version"`
	Grpc    KratosGrpc `json:"Grpc" yaml:"Grpc" mapstructure:"Grpc"`
	Http    KratosHttp `json:"Http" yaml:"Http" mapstructure:"Http"`
}

type KratosGrpc struct {
	Ip      string        `json:"Ip" yaml:"Ip" mapstructure:"Ip"`
	Port    string        `json:"Port" yaml:"Port" mapstructure:"Port"`
	Timeout time.Duration `json:"Timeout" yaml:"Timeout" mapstructure:"Timeout"`
}

func (c KratosGrpc) Target() string {
	return fmt.Sprintf("%s:%s", c.Ip, c.Port)
}

type KratosHttp struct {
	Ip      string        `json:"Ip" yaml:"Ip" mapstructure:"Ip"`
	Port    string        `json:"Port" yaml:"Port" mapstructure:"Port"`
	Timeout time.Duration `json:"Timeout" yaml:"Timeout" mapstructure:"Timeout"`
}

func (c KratosHttp) Address() string {
	return fmt.Sprintf("%s:%s", c.Ip, c.Port)
}
