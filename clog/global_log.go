package clog

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog"
)

var (
	global      ContextLogger
	globalLevel zerolog.Level
)

func init() {

	zerolog.LevelTraceValue = "TRACE"
	zerolog.LevelDebugValue = "DEBUG"
	zerolog.LevelInfoValue = "INFO"
	zerolog.LevelWarnValue = "WARN"
	zerolog.LevelErrorValue = "ERROR"
	zerolog.LevelFatalValue = "FATAL"
	zerolog.LevelPanicValue = "PANIC"

	global = newGlobalLogger()
}

func SetGlobalLevel(level zerolog.Level) {
	globalLevel = level
	zerolog.SetGlobalLevel(level)
}
func newGlobalLogger() ContextLogger {
	// グローバルロガーの場合、呼び出し元を正しく表示するために適切なCallSkip値を設定
	out := NewCustomConsoleWriter(2)

	// CallerWithSkipFrameCountは使用せず、CustomConsoleWriterのCallSkipのみで制御
	l := zerolog.New(out).With().Logger()
	impl := implContextLogger{
		logger: l,
	}

	return &impl
}

func SetGlobalLog(logger ContextLogger) {
	global = logger
}

func GlobalLogger() ContextLogger {
	return global
}

func Debug(i ...interface{}) {
	global.Debug(i...)
}

func Debugf(format string, args ...interface{}) {
	global.Debugf(format, args...)
}

func Debugj(j JSON) {
	global.Debug(toJson(j))
}

func Info(i ...interface{}) {
	global.Info(i...)
}

func Infof(format string, args ...interface{}) {
	global.Infof(format, args...)
}

func Infoj(j JSON) {
	global.Info(toJson(j))
}

func Warn(i ...interface{}) {
	global.Warn(i...)
}

func Warnf(format string, args ...interface{}) {
	global.Warnf(format, args...)
}

func Warnj(j JSON) {
	global.Warn(toJson(j))
}

func Error(i ...interface{}) {
	global.Error(i...)
}

func Errorf(format string, args ...interface{}) {
	global.Errorf(format, args...)
}

func Errorj(j JSON) {
	global.Error(toJson(j))
}

func Fatal(i ...interface{}) {
	global.Fatal(i...)
}

func Fatalf(format string, args ...interface{}) {
	global.Fatalf(format, args...)
}

func Fatalj(j JSON) {
	global.Fatal(toJson(j))
}

func Panic(i ...interface{}) {
	global.Panic(i...)
}

func Panicf(format string, args ...interface{}) {
	global.Panicf(format, args...)
}

func Panicj(j JSON) {
	global.Panic(toJson(j))
}

func Ctx(ctx context.Context) ContextLogger {
	logger, ok := ctx.Value(LOGGING_CONTECT_KEY).(ContextLogger)
	if !ok {
		global.Warn(`contextにlogインスタンスがありません`)
		return global
	}

	return logger

}

const LOGGING_CONTECT_KEY = `__contextLog__`

func toJson(j JSON) string {
	if j == nil {
		return "nil"
	}

	jsonAsByte, _ := json.Marshal(j)
	return string(jsonAsByte)
}
