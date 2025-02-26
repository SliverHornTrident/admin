//go:build (gorm || gorms) && (tidb || mysql || postgres || sqlite || clickhouse || mssql || sqlserver || oracle)

package config

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Gorm struct {
	Host            string         `json:"Host" yaml:"Host" mapstructure:"Host"`                                  // 数据库地址
	Port            int64          `json:"Port" yaml:"Port" mapstructure:"Port"`                                  // 数据库端口
	Config          string         `json:"Config" yaml:"Config" mapstructure:"Config"`                            // 数据库其他配置
	Dbname          string         `json:"Dbname" yaml:"Dbname" mapstructure:"Dbname"`                            // 数据库名
	Username        string         `json:"Username" yaml:"Username" mapstructure:"Username"`                      // 数据库用户名
	Password        string         `json:"Password" yaml:"Password" mapstructure:"Password"`                      // 数据库密码
	ConnMaxLifetime time.Duration  `json:"ConnMaxLifetime" yaml:"ConnMaxLifetime" mapstructure:"ConnMaxLifetime"` // 连接最大存活时间
	ConnMaxIdleTime time.Duration  `json:"ConnMaxIdleTime" yaml:"ConnMaxIdleTime" mapstructure:"ConnMaxIdleTime"` // 连接最大空闲时间
	MaxIdleCones    int            `json:"MaxIdleCones" yaml:"MaxIdleCones" mapstructure:"MaxIdleCones"`          // 最大空闲连接数
	MaxOpenCones    int            `json:"MaxOpenCones" yaml:"MaxOpenCones" mapstructure:"MaxOpenCones"`          // 最大打开连接数
	Logger          GormLogger     `json:"Logger" yaml:"Logger" mapstructure:"Logger"`                            // 日志配置
	Resolvers       []GormResolver `json:"Resolvers" yaml:"Resolvers" mapstructure:"Resolvers"`                   // 数据库负载均衡
}

func (c Gorm) IsEmpty() bool {
	if c.Host == "" || c.Port == 0 || c.Dbname == "" || c.Username == "" || c.Password == "" {
		return true
	}
	return false
}

type GormResolver struct {
	Sources           []Gorm        `json:"Sources" yaml:"Sources" mapstructure:"Sources"`                               // 主库配置
	Replicas          []Gorm        `json:"Replicas" yaml:"Replicas" mapstructure:"Replicas"`                            // 从库配置
	ConnMaxLifetime   time.Duration `json:"ConnMaxLifetime" yaml:"ConnMaxLifetime" mapstructure:"ConnMaxLifetime"`       // 连接最大存活时间
	ConnMaxIdleTime   time.Duration `json:"ConnMaxIdleTime" yaml:"ConnMaxIdleTime" mapstructure:"ConnMaxIdleTime"`       // 连接最大空闲时间
	MaxIdleCones      int           `json:"MaxIdleCones" yaml:"MaxIdleCones" mapstructure:"MaxIdleCones"`                // 最大空闲连接数
	MaxOpenCones      int           `json:"MaxOpenCones" yaml:"MaxOpenCones" mapstructure:"MaxOpenCones"`                // 最大打开连接数
	TraceResolverMode bool          `json:"TraceResolverMode" yaml:"TraceResolverMode" mapstructure:"TraceResolverMode"` // 是否开启负载均衡日志跟踪模式
	Datasets          []string      `json:"Datasets" yaml:"Datasets" mapstructure:"Datasets"`                            // 数据表名
}

func (c GormResolver) Data() []any {
	length := len(c.Datasets)
	data := make([]any, 0, length)
	for i := 0; i < length; i++ {
		data = append(data, c.Datasets[i])
	}
	return data
}

func (c Gorm) GormConfig() *gorm.Config {
	slowThreshold, err := time.ParseDuration(c.Logger.SlowThreshold)
	if err != nil {
		slowThreshold = 200 * time.Millisecond // 默认200ms
	}
	config := logger.Config{
		LogLevel:                  c.Logger.LogLevel,
		Colorful:                  c.Logger.Colorful,
		SlowThreshold:             slowThreshold,
		IgnoreRecordNotFoundError: c.Logger.IgnoreRecordNotFoundError,
	}
	writer := NewWriter(c, log.New(os.Stdout, "\r\n", log.LstdFlags))
	return &gorm.Config{
		Logger:                                   logger.New(writer, config),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

type GormLogger struct {
	Zap                       bool            `json:"Zap" yaml:"Zap" mapstructure:"Zap"`                                                                   // 日志是否zap持久化
	Console                   bool            `json:"Console" yaml:"Console" mapstructure:"Console"`                                                       // 是否打印到控制台
	Colorful                  bool            `json:"Colorful" yaml:"Colorful" mapstructure:"Colorful"`                                                    // 是否彩色
	LogLevel                  logger.LogLevel `json:"LogLevel" yaml:"LogLevel" mapstructure:"LogLevel"`                                                    // 日志级别(1:silent,2:error,3:warn,4:info)
	SlowThreshold             string          `json:"SlowThreshold" yaml:"SlowThreshold" mapstructure:"SlowThreshold"`                                     // 慢查询阈值(支持秒[s],分[m],时[h])
	IgnoreRecordNotFoundError bool            `json:"IgnoreRecordNotFoundError" yaml:"IgnoreRecordNotFoundError" mapstructure:"IgnoreRecordNotFoundError"` // 忽略记录未找到错误
}
