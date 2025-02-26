//go:build mongo

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"github.com/qiniu/qmgo"
)

var (
	Mongo       *qmgo.QmgoClient
	MongoConfig config.Mongo
)
