//go:build zap

package config

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type Zap struct {
	Level        string `json:"Level" yaml:"Level" mapstructure:"Level"`                      // 级别
	Prefix       string `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"`                   // 日志前缀
	Format       string `json:"Format" yaml:"Format" mapstructure:"Format"`                   // 输出
	Director     string `json:"Director" yaml:"Director" mapstructure:"Director"`             // 日志文件夹
	ShowLine     bool   `json:"ShowLine" yaml:"ShowLine" mapstructure:"ShowLine"`             // 显示行
	InConsole    bool   `json:"InConsole" yaml:"InConsole" mapstructure:"InConsole"`          // 在控制台显示
	InOutputFile bool   `json:"InOutputFile" yaml:"InOutputFile" mapstructure:"InOutputFile"` // 写入文件中显示
}

// Levels 根据字符串转化为 zapcore.Levels
func (c *Zap) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}

func (c *Zap) Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(c.Prefix + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if c.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}
