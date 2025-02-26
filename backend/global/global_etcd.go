package global

import (
	"github.com/SliverHornTrident/shadow/config"
	etcd "go.etcd.io/etcd/client/v3"
)

var (
	Etcd       *etcd.Client
	EtcdConfig *config.Etcd
)
