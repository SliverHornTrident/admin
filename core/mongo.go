//go:build mongo

package core

import (
	"context"
	"fmt"
	"github.com/SliverHornTrident/shadow/core/internal"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/SliverHornTrident/shadow/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	option "go.mongodb.org/mongo-driver/mongo/options"
	"sort"
	"strings"
)

var _ interfaces.Corer = (*mongo)(nil)

var Mongo = new(mongo)

type mongo struct{}

func (c *mongo) Name() string {
	return "[shadow][core][mongo]"
}

func (c *mongo) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Mongo", &global.MongoConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *mongo) IsPanic() bool {
	return true
}

func (c *mongo) ConfigName() string {
	return strings.Join([]string{"mongo", gin.Mode(), "yaml"}, ".")
}

func (c *mongo) Initialization(ctx context.Context) error {
	var opts []options.ClientOptions
	if global.MongoConfig.IsZap {
		opts = internal.Mongo.GetClientOptions()
	}
	client, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:              global.MongoConfig.Uri(),
		Coll:             global.MongoConfig.Coll,
		Database:         global.MongoConfig.Database,
		MinPoolSize:      &global.MongoConfig.MinPoolSize,
		MaxPoolSize:      &global.MongoConfig.MaxPoolSize,
		SocketTimeoutMS:  &global.MongoConfig.SocketTimeoutMS,
		ConnectTimeoutMS: &global.MongoConfig.ConnectTimeoutMS,
		Auth: &qmgo.Credential{
			Username: global.MongoConfig.Username,
			Password: global.MongoConfig.Password,
		},
	}, opts...)
	if err != nil {
		return errors.Wrap(err, "链接mongodb数据库失败!")
	}
	global.Mongo = client
	err = c.Indexes(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *mongo) Indexes(ctx context.Context) error {
	indexMap := map[string][][]string{}
	for collection, indexes := range indexMap {
		err := c.CreateIndexes(ctx, collection, indexes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *mongo) CreateIndexes(ctx context.Context, name string, indexes [][]string) error {
	collection, err := global.Mongo.Database.Collection(name).CloneCollection()
	if err != nil {
		global.Mongo.Database.Collection(name).InsertOne(ctx, bson.M{})
		return errors.Wrapf(err, "获取[%s]的表对象失败!", name)
	}
	collection.InsertOne(ctx, bson.M{})
	list, err := collection.Indexes().List(ctx)
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的索引对象失败!", name)
	}
	var entities []Index
	err = list.All(ctx, &entities)
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的索引列表失败!", name)
	}
	length := len(indexes)
	indexMap1 := make(map[string][]string, length)
	for i := 0; i < length; i++ {
		sort.Strings(indexes[i])
		length1 := len(indexes[i])
		keys := make([]string, 0, length1)
		for j := 0; j < length1; j++ {
			if indexes[i][i][0] == '-' {
				keys = append(keys, indexes[i][j], "-1")
				continue
			}
			keys = append(keys, indexes[i][j], "1")
		}
		key := strings.Join(keys, "_")
		_, o1 := indexMap1[key]
		if o1 {
			return errors.Errorf("索引[%s]重复!", key)
		}
		indexMap1[key] = indexes[i]
	}
	length = len(entities)
	indexMap2 := make(map[string]map[string]string, length)
	for i := 0; i < length; i++ {
		v1, o1 := indexMap2[entities[i].Name]
		if !o1 {
			keyLength := len(entities[i].Key)
			v1 = make(map[string]string, keyLength)
			for j := 0; j < keyLength; j++ {
				v2, o2 := v1[entities[i].Key[j].Key]
				if !o2 {
					v1 = make(map[string]string)
				}
				v2 = entities[i].Key[j].Key
				v1[entities[i].Key[j].Key] = v2
				indexMap2[entities[i].Name] = v1
			}
		}
	}
	for k1, v1 := range indexMap1 {
		_, o2 := indexMap2[k1]
		if o2 {
			continue
		} // 索引存在
		if len(fmt.Sprintf("%s.%s.$%s", collection.Name(), name, v1)) > 127 {
			err = global.Mongo.Database.Collection(name).CreateOneIndex(ctx, options.IndexModel{
				Key:          v1,
				IndexOptions: option.Index().SetName(utils.Md5(k1)).SetExpireAfterSeconds(86400),
			})
			if err != nil {
				return errors.Wrapf(err, "创建索引[%s]失败!", k1)
			}
			return nil
		}
		err = global.Mongo.Database.Collection(name).CreateOneIndex(ctx, options.IndexModel{
			Key:          v1,
			IndexOptions: option.Index().SetExpireAfterSeconds(86400),
		})
		if err != nil {
			return errors.Wrapf(err, "创建索引[%s]失败!", k1)
		}
	}
	return nil
}

type Index struct {
	V    any      `bson:"v"`
	Ns   any      `bson:"ns"`
	Key  []bson.E `bson:"key"`
	Name string   `bson:"name"`
}
