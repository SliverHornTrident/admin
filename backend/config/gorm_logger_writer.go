//go:build (gorm || gorms) && (tidb || mysql || postgres || sqlite || clickhouse || mssql || sqlserver || oracle)

package config

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Writer struct {
	config Gorm
	writer logger.Writer
}

func NewWriter(config Gorm, writer logger.Writer) *Writer {
	return &Writer{config: config, writer: writer}
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...any) {
	if c.config.Logger.Console {
		c.writer.Printf(message, data...)
		if !c.config.Logger.Zap {
			return
		} // 如果不是zap持久化，直接return
	} // 控制台输出
	switch c.config.Logger.LogLevel {
	case logger.Silent:
		zap.L().Debug(fmt.Sprintf(message, data...))
	case logger.Error:
		zap.L().Error(fmt.Sprintf(message, data...))
	case logger.Warn:
		zap.L().Warn(fmt.Sprintf(message, data...))
	case logger.Info:
		zap.L().Info(fmt.Sprintf(message, data...))
	default:
		zap.L().Info(fmt.Sprintf(message, data...))
	}
	return
}
