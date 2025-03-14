package clog

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

func NewContextLogger() ContextLogger {
	// コンテキストロガーの場合、呼び出し元を正しく表示するために適切なCallSkip値を設定
	out := NewCustomConsoleWriter(6)

	// CallerWithSkipFrameCountは使用せず、CustomConsoleWriterのCallSkipのみで制御
	l := zerolog.New(out).
		With().
		Logger().
		Level(globalLevel)

	impl := implContextLogger{
		logger: l,
	}

	return &impl
}

func (this *implContextLogger) Debugj(j JSON) {
	this.logger.Debug().Msg(toJson(j))
}

func (this *implContextLogger) Infoj(j JSON) {
	this.logger.Info().Msg(toJson(j))
}

func (this *implContextLogger) Warnj(j JSON) {
	this.logger.Warn().Msg(toJson(j))
}

func (this *implContextLogger) Errorj(j JSON) {
	this.logger.Error().Msg(toJson(j))
}

func (this *implContextLogger) Fatalj(j JSON) {
	this.logger.Fatal().Msg(toJson(j))
}

func (this *implContextLogger) Panicj(j JSON) {
	this.logger.Debug().Msg(toJson(j))
}

type implContextLogger struct {
	logger zerolog.Logger
	prefix map[string]interface{}
}

func (this *implContextLogger) SetPrefix(key string, val interface{}) {

	if this.prefix == nil {
		this.prefix = map[string]interface{}{}
	}
	this.prefix[key] = val

	this.logger = this.logger.With().Any(`prefix`, this.prefix).Logger()

}

func ifc(i ...interface{}) string {
	s := []string{}
	for _, ii := range i {
		s = append(s, fmt.Sprint(ii))
	}
	return strings.Join(s, ` `)
}

func (this *implContextLogger) Debug(i ...interface{}) {
	this.logger.Debug().Msg(ifc(i...))
}

func (this *implContextLogger) Debugf(format string, args ...interface{}) {
	this.logger.Debug().Msgf(format, args...)
}

func (this *implContextLogger) Info(i ...interface{}) {
	this.logger.Info().Msg(ifc(i...))
}

func (this *implContextLogger) Infof(format string, args ...interface{}) {
	this.logger.Info().Msgf(format, args...)
}

func (this *implContextLogger) Warn(i ...interface{}) {
	this.logger.Warn().Msg(ifc(i...))
}

func (this *implContextLogger) Warnf(format string, args ...interface{}) {
	this.logger.Warn().Msgf(format, args...)
}

func (this *implContextLogger) Error(i ...interface{}) {
	this.logger.Error().Msg(ifc(i...))
}

func (this *implContextLogger) Errorf(format string, args ...interface{}) {
	this.logger.Error().Msgf(format, args...)
}

func (this *implContextLogger) Fatal(i ...interface{}) {
	this.logger.Fatal().Msg(ifc(i...))
}

func (this *implContextLogger) Fatalf(format string, args ...interface{}) {
	this.logger.Fatal().Msgf(format, args...)
}

func (this *implContextLogger) Panic(i ...interface{}) {
	this.logger.Panic().Msg(ifc(i...))
}

func (this *implContextLogger) Panicf(format string, args ...interface{}) {
	this.logger.Panic().Msgf(format, args...)
}

func (this *implContextLogger) WithContext(oldContext context.Context) (newContext context.Context) {
	newContext = context.WithValue(oldContext, LOGGING_CONTECT_KEY, this)
	return newContext
}
