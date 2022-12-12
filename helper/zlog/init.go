package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func InitLogger(useProduction bool) {
	var cfg zap.Config

	if useProduction {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	z, _ := cfg.Build()
	logger = z.WithOptions(zap.WithCaller(false)).Sugar()
}
