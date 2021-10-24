package logger

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger interface {
	Debug(requiredId string, args ...interface{})
	Info(requiredId string, args ...interface{})
	Warn(requiredId string, args ...interface{})
	Error(requiredId string, args ...interface{})
}
