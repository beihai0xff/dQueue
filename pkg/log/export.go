package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Debug logs to DEBUG log
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info logs to INFO log
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn logs to WARN log
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error logs to ERROR log
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal logs to FATAL log
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Debugf logs to DEBUG log, Arguments will be formatted with Sprint, Sprintf, or neither.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Infof logs to INFO log, Arguments will be formatted with Sprint, Sprintf, or neither.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warnf logs to WARN log, Arguments will be formatted with Sprint, Sprintf, or neither.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Errorf logs to ERROR log, Arguments will be formatted with Sprint, Sprintf, or neither.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatalf logs to FATAL log, Arguments will be formatted with Sprint, Sprintf, or neither.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Sync calls the zap defaultLogger's Sync method, flushing any buffered log entries.
// Applications should take care to call Sync before exiting.
func Sync() error {
	return defaultLogger.Sync()
}

// WithFields 设置一些业务自定义数据到每条 log 中
func WithFields(fields ...interface{}) *zap.SugaredLogger {
	return defaultLogger.WithOptions(zap.AddStacktrace(zapcore.WarnLevel)).With(fields...)
}
