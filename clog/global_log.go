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
	// スタックトレースの開始位置を設定（実際の呼び出し元は自動検出される）
	out := NewCustomConsoleWriter(3)

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

// Clone はコンテキストからロガーを取得し、そのコピーを作成して新しいコンテキストとロガーを返します。
func Clone(ctx context.Context) (newContext context.Context, newLogger ContextLogger) {
	// コンテキストからロガーを取得
	logger, ok := ctx.Value(LOGGING_CONTECT_KEY).(*implContextLogger)
	if !ok {
		global.Warn(`contextにlogインスタンスがありません`)
		// グローバルロガーのコピーを作成
		newImpl := &implContextLogger{
			logger: global.(*implContextLogger).logger,
		}
		newContext = context.WithValue(ctx, LOGGING_CONTECT_KEY, newImpl)
		newLogger = newImpl
		return newContext, newLogger
	}

	// ロガーのコピーを作成
	newImpl := &implContextLogger{
		logger: logger.logger,
	}
	
	// prefixがある場合はコピー
	if logger.prefix != nil {
		newImpl.prefix = make(map[string]interface{})
		for k, v := range logger.prefix {
			newImpl.prefix[k] = v
		}
	}
	
	// 元のコンテキストをベースに新しいコンテキストを作成
	newContext = context.WithValue(ctx, LOGGING_CONTECT_KEY, newImpl)
	newLogger = newImpl
	
	return newContext, newLogger
}

const LOGGING_CONTECT_KEY = `__contextLog__`

func toJson(j JSON) string {
	if j == nil {
		return "nil"
	}

	jsonAsByte, _ := json.Marshal(j)
	return string(jsonAsByte)
}
