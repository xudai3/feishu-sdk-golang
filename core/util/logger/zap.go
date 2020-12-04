package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.SugaredLogger

func InitLogger(DebugLevel string) (err error) {
	writerSyncer := os.Stdout
	encoder := getEncoder()
	// 开启文件和行号
	caller := zap.AddCaller()
	var l zapcore.Level
	err = l.Set(DebugLevel)
	if err != nil {
		fmt.Printf("zap logger set log level failed:%v\n", err)
		return err
	}
	core := zapcore.NewCore(encoder, writerSyncer, l)
	logger = zap.New(core, caller, zap.AddCallerSkip(1)).Sugar()
	defer logger.Sync()
	return nil
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(args ... interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
}
