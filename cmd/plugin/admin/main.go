package main

import (
	"github.com/SliverHornTrident/shadow/core"
	admin "github.com/SliverHornTrident/shadow/plugin/admin/core"
)

func main() {
	core.Register(core.Viper, core.Zap, core.Mysql, admin.Gorm, core.Gin)
}
