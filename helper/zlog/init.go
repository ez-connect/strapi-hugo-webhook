package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"strapiwebhook/service/config"
)

var logger *zap.SugaredLogger

func InitLogger() {
	var cfg zap.Config

	if config.BuildMode == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	z, _ := cfg.Build()
	logger = z.WithOptions(zap.WithCaller(false)).Sugar()
}
