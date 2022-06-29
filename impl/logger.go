package impl

import (
	"go.uber.org/zap"

	"strapiwebhook/base"
)

type Logger struct {
	*zap.SugaredLogger
}

// go-kit log interface used for metrics
// func New(prefix string, logger log.Logger) *Graphite {
func (l Logger) Log(keyvals ...interface{}) error {
	l.Infow("", keyvals...)
	return nil
}

// type WriteSyncer struct {
// 	io.Writer
// }

var logger Logger

func GetLogger() Logger {
	return logger
}

func InitLogger() {
	var z *zap.Logger

	if base.BuildMode == "production" {
		z, _ = zap.NewProduction()
	} else {
		z, _ = zap.NewDevelopment()
	}

	logger.SugaredLogger = z.Sugar()

	// // logger into file
	// // logpath := "./sample.log"
	// // if runtime.GOOS == "linux" {
	// //	logpath = "/var/log/sample.log"
	// // }
	// syncer := zap.CombineWriteSyncers(os.Stdout /*, getWriteSyncer(logpath)*/)
	// pe := zap.NewProductionEncoderConfig()
	// fileEncoder := zapcore.NewJSONEncoder(pe)
	// core := zapcore.NewCore(fileEncoder, syncer, zap.NewAtomicLevelAt(zap.InfoLevel))
	// logger.SugaredLogger = zap.New(core).WithOptions(zap.AddCaller()).Sugar()
}

// func (ws WriteSyncer) Sync() error {
// 	return nil
// }

// func getWriteSyncer(logName string) zapcore.WriteSyncer {
// 	var ioWriter = &lumberjack.Logger{
// 		Filename:   logName,
// 		MaxSize:    10, // MB
// 		MaxBackups: 3,  // number of backups
// 		MaxAge:     28, //days
// 		LocalTime:  true,
// 		Compress:   false, // disabled by default
// 	}
// 	var sw = WriteSyncer{ioWriter}
// 	return sw
// }

