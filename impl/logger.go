package impl

import (
	"fmt"
	"io"
	"os"

	// "runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.Logger
}
type WriteSyncer struct {
	io.Writer
}

func (ws WriteSyncer) Sync() error {
	return nil
}

var logger = Logger{}

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

func GetLogger() Logger {
	if logger.Logger == nil {
		// logger into file
		// logpath := "./sample.log"
		// if runtime.GOOS == "linux" {
		//	logpath = "/var/log/sample.log"
		// }
		syncer := zap.CombineWriteSyncers(os.Stdout /*, getWriteSyncer(logpath)*/)
		pe := zap.NewProductionEncoderConfig()
		fileEncoder := zapcore.NewJSONEncoder(pe)
		core := zapcore.NewCore(fileEncoder, syncer, zap.NewAtomicLevelAt(zap.InfoLevel))
		logger.Logger = zap.New(core).WithOptions(zap.AddCaller())
	}
	return logger
}

func (l Logger) Log(keyvals ...interface{}) error {
	l.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprint(keyvals...))
	return nil
}
