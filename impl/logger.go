package impl

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"strapiwebhook/base"
)

type Logger struct {
	*zap.SugaredLogger
}

// Implement `go-kit`` log interface used for metrics
// func New(prefix string, logger log.Logger) *Graphite {
func (l Logger) Log(keyvals ...interface{}) error {
	l.Infow("", keyvals...)
	return nil
}

var logger Logger

func GetLogger() Logger {
	return logger
}

func InitLogger() {
	if logger.SugaredLogger != nil {
		return
	}

	var cfg zap.Config

	if base.BuildMode == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	z, _ := cfg.Build()
	logger.SugaredLogger = z.WithOptions(zap.WithCaller(false)).Sugar()
}
