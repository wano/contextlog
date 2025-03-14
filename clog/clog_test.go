package clog

import (
	"context"
	"runtime/debug"
	"testing"

	"golang.org/x/xerrors"
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

func TestClone(t *testing.T) {
	// 元のロガーを作成
	originalLogger := NewContextLogger()
	originalLogger.SetPrefix("original", "value")
	
	// コンテキストに関連付け
	ctx := context.Background()
	ctx = originalLogger.WithContext(ctx)
	
	// Clone メソッドを使用して新しいコンテキストとロガーを取得
	newCtx, newLogger := Clone(ctx)
	
	// 新しいロガーにプレフィックスを設定
	newLogger.SetPrefix("new", "value")
	
	// 元のコンテキストからロガーを取得
	loggerFromOriginalCtx := Ctx(ctx)
	
	// 新しいコンテキストからロガーを取得
	loggerFromNewCtx := Ctx(newCtx)
	
	// テスト: 元のコンテキストと新しいコンテキストから取得したロガーが異なること
	if loggerFromOriginalCtx == loggerFromNewCtx {
		t.Error("Clone should create a new logger instance")
	}
	
	// ログ出力のテスト
	loggerFromOriginalCtx.Info("Log from original context")
	loggerFromNewCtx.Info("Log from new context")
}
