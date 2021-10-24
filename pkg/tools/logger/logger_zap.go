package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	sugar *zap.SugaredLogger
}

func NewLoggerZap(levelOfLogging LogLevel) Logger {
	var loggingLevel zapcore.Level
	switch levelOfLogging {
	case DebugLevel:
		loggingLevel = zap.DebugLevel
	case InfoLevel:
		loggingLevel = zap.InfoLevel
	case WarnLevel:
		loggingLevel = zap.WarnLevel
	case ErrorLevel:
		loggingLevel = zap.ErrorLevel
	default:
		loggingLevel = zap.DebugLevel
	}

	fileName := "var/log/zap.log"
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30,
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter,
		zap.NewAtomicLevelAt(loggingLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &LoggerZap{
		sugar: logger.Sugar(),
	}
}

func (l *LoggerZap) Debug(requiredId string, args ...interface{}) {
	l.sugar.Debugw(requiredId, args...)
}

func (l *LoggerZap) Info(requiredId string, args ...interface{}) {
	l.sugar.Infow(requiredId, args...)
}

func (l *LoggerZap) Warn(requiredId string, args ...interface{}) {
	l.sugar.Warnw(requiredId, args...)
}

func (l *LoggerZap) Error(requiredId string, args ...interface{}) {
	l.sugar.Errorw(requiredId, args...)
}
