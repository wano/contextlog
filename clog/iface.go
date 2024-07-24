package clog

import "context"

type ContextLogger interface {
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})
	Debugj(j JSON)
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Infoj(j JSON)
	Warn(i ...interface{})
	Warnf(format string, args ...interface{})
	Warnj(j JSON)
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
	Errorj(j JSON)
	Fatal(i ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalj(j JSON)
	Panic(i ...interface{})
	Panicf(format string, args ...interface{})
	Panicj(j JSON)
	SetPrefix(s string, val interface{})
	WithContext(oldContext context.Context) (newContext context.Context)
}
type (
	JSON = map[string]interface{}
)
