package clog

import (
	"context"
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
