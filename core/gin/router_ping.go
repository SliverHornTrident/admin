//go:build gin

package core

import (
	"github.com/SliverHornTrident/shadow/constant"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

var RouterPing = new(routerPing)

type routerPing struct{}

func (e *routerPing) Initialization(public *gin.RouterGroup) {
	public.GET("ping", func(c *gin.Context) {
		ip := global.GinConfig.Host
		if ip == "" {
			address, _ := net.InterfaceAddrs()
			for i := 0; i < len(address); i++ { // 这个网络地址是IP地址: ipv4, ipv6
				ipNet, isIpNet := address[i].(*net.IPNet)
				if isIpNet && !ipNet.IP.IsLoopback() { // 跳过IPV6
					if ipNet.IP.To4() != nil {
						ip = ipNet.IP.String()
						break
					}
				}
			}
		}
		data := struct {
			Ip       string `json:"ip"`
			Port     int    `json:"port"`
			Name     string `json:"name"`
			Node     string `json:"node"`
			Message  string `json:"message"`
			Version  string `json:"version"`
			ClientIp string `json:"client_ip"`
		}{
			Ip:       ip,
			Port:     global.GinConfig.Port,
			Name:     global.GinConfig.Name,
			Node:     global.GinConfig.Node,
			Message:  "pong",
			Version:  constant.Version,
			ClientIp: c.ClientIP(),
		}
		c.JSON(http.StatusOK, data)
	})
}
