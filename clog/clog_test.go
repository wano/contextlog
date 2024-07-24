package clog

import (
	"context"
	"golang.org/x/xerrors"
	"runtime/debug"
	"testing"
)

func TestCallStack(t *testing.T) {

	c := NewContextLogger()
	c.SetPrefix("key", "value")

	ctx := context.Background()
	ctx = c.WithContext(ctx)

	Ctx(ctx).Info("Context-LOOG")

	Info("GlobalLog")

}

func TestCallErrStack(t *testing.T) {

	c := NewContextLogger()

	defer func() {
		if onPanic := recover(); onPanic != nil {
			err := xerrors.Errorf("recover: %v \n %v", onPanic, string(debug.Stack()))
			c.Error(err)
		}
	}()

	ctx := context.Background()
	ctx = c.WithContext(ctx)

	panic("Panic")
}
