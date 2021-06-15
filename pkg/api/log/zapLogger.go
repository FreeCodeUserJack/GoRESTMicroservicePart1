package log

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ZapLogger *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding: "json",
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey: "level",
			TimeKey: "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	ZapLogger, err = logConfig.Build()

	if err != nil {
		panic(err)
	}
}

func DebugZap(msg string, tags ...string) {
	ZapLogger.Debug(msg, getZapFields(tags...)...)
	ZapLogger.Sync()
}

func InfoZap(msg string, tags ...string) {
	ZapLogger.Info(msg, getZapFields(tags...)...)
	ZapLogger.Sync()
}

func ErrorZap(msg string, err error, tags ...string) {
	msg = fmt.Sprintf("%s - ERROR = %V", msg, err)
	ZapLogger.Error(msg, getZapFields(tags...)...)
	ZapLogger.Sync()
}

func getZapFields(tags ...string) []zap.Field {
	result := make([]zap.Field, len(tags))

	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result = append(result, zap.Any(strings.TrimSpace(els[0]), strings.TrimSpace(els[1])))
	}

	return result
}