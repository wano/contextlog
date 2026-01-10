# contextlog

[github.com/labstack/gommon/log](https://github.com/labstack/gommon/tree/master/log) のインターフェースを実装した [zerolog](https://github.com/rs/zerolog) ベースの構造化ロガーです。

## 特徴

- **gommon/log 互換インターフェース**: labstack/gommon/log と同じ API を提供
- **構造化ログ**: JSON 形式でのログ出力に対応
- **コンテキスト対応**: Go の context を利用したリクエストスコープのログ管理
- **自動呼び出し元検出**: ログ出力箇所のファイル名・行番号を自動付与
- **プレフィックス機能**: リクエストIDやアプリ名などのメタデータを全ログに付与
- **スレッドセーフ**: zerolog による安全な並行処理対応
- **ロガーのクローン機能**: コンテキストから独立したロガーインスタンスを生成可能

## なぜ contextlog を使うのか

- `labstack/gommon/log` はカスタマイズ性に乏しい
- `labstack/gommon/log` の開発が止まっている
- スレッドロック等、zap 以降のロガーを使っていた方が[良さげ](https://zenn.dev/moriyoshi/articles/1af0659e29d727)

## インストール

```bash
go get -u github.com/wano/contextlog/...
```

## 使い方

### グローバルロガー

最もシンプルな使い方です。グローバルロガーを使用してログを出力します。

```go
import "github.com/wano/contextlog/clog"

// 通常のログ出力
clog.Debug("デバッグメッセージ")
clog.Info("情報メッセージ")
clog.Warn("警告メッセージ")
clog.Error("エラーメッセージ")

// フォーマット付きログ出力
clog.Infof("ユーザー %s がログインしました", username)

// JSON構造化ログ出力
clog.Infoj(map[string]interface{}{
    "user_id": 123,
    "action":  "login",
})
```

### ログレベルの制御

```go
import "github.com/rs/zerolog"

// グローバルログレベルを設定
clog.SetGlobalLevel(zerolog.ErrorLevel)

clog.Info("この出力は表示されません")
clog.Error("この出力は表示されます")
```

### コンテキスト対応ロガー

リクエストスコープでログにメタデータを付与する場合に使用します。

```go
import (
    "context"
    "github.com/wano/contextlog/clog"
)

// ロガーを作成してコンテキストに紐付け
ctx := context.Background()
logger := clog.NewContextLogger()
logger.SetPrefix("app_name", "my-service")
logger.SetPrefix("request_id", "req-12345")
ctx = logger.WithContext(ctx)

// コンテキストからロガーを取得してログ出力
clog.Ctx(ctx).Info("リクエストを処理中")
clog.Ctx(ctx).Infof("ユーザー %s の処理完了", username)
```

出力例:
```json
{
  "level": "INFO",
  "msg": "リクエストを処理中",
  "meta": {
    "app_name": "my-service",
    "request_id": "req-12345"
  },
  "timestamp": "2024-01-10T15:04:05.000Z",
  "caller": "/path/to/handler.go:42"
}
```

### ロガーのクローン

既存のコンテキストから独立したロガーインスタンスを生成できます。元のロガーのプレフィックスを継承しつつ、新しいプレフィックスを追加しても元のロガーに影響しません。

```go
// 元のコンテキストとロガー
ctx := context.Background()
logger := clog.NewContextLogger()
logger.SetPrefix("app_name", "my-service")
ctx = logger.WithContext(ctx)

// クローンを作成（プレフィックスを継承）
newCtx, newLogger := clog.Clone(ctx)
newLogger.SetPrefix("sub_task", "task-001")  // 元のロガーには影響しない

clog.Ctx(ctx).Info("元のロガー")       // app_name のみ
clog.Ctx(newCtx).Info("クローン")       // app_name + sub_task
```

## インターフェース

`ContextLogger` インターフェースは以下のメソッドを提供します:

| メソッド | 説明 |
|---------|------|
| `Debug(i ...interface{})` | デバッグレベルのログ |
| `Debugf(format string, i ...interface{})` | フォーマット付きデバッグログ |
| `Debugj(j JSON)` | JSON構造化デバッグログ |
| `Info(i ...interface{})` | 情報レベルのログ |
| `Infof(format string, i ...interface{})` | フォーマット付き情報ログ |
| `Infoj(j JSON)` | JSON構造化情報ログ |
| `Warn(i ...interface{})` | 警告レベルのログ |
| `Warnf(format string, i ...interface{})` | フォーマット付き警告ログ |
| `Warnj(j JSON)` | JSON構造化警告ログ |
| `Error(i ...interface{})` | エラーレベルのログ |
| `Errorf(format string, i ...interface{})` | フォーマット付きエラーログ |
| `Errorj(j JSON)` | JSON構造化エラーログ |
| `Fatal(i ...interface{})` | 致命的エラーログ（プログラム終了） |
| `Fatalf(format string, i ...interface{})` | フォーマット付き致命的エラーログ |
| `Fatalj(j JSON)` | JSON構造化致命的エラーログ |
| `Panic(i ...interface{})` | パニックログ（panic発生） |
| `Panicf(format string, i ...interface{})` | フォーマット付きパニックログ |
| `Panicj(j JSON)` | JSON構造化パニックログ |
| `SetPrefix(key, value string)` | プレフィックス（メタデータ）を設定 |
| `WithContext(ctx context.Context) context.Context` | ロガーをコンテキストに紐付け |

詳細は [interface](./clog/iface.go) を参照してください。

## グローバル関数

| 関数 | 説明 |
|------|------|
| `GlobalLogger()` | グローバルロガーインスタンスを取得 |
| `SetGlobalLevel(level zerolog.Level)` | グローバルログレベルを設定 |
| `Ctx(ctx context.Context) ContextLogger` | コンテキストからロガーを取得（なければグローバルロガー） |
| `Clone(ctx context.Context) (context.Context, ContextLogger)` | コンテキストからロガーをクローン |

## 出力フォーマット

ログは以下のJSON形式で出力されます:

```json
{
  "level": "INFO",
  "msg": "メッセージ内容",
  "meta": {
    "key1": "value1",
    "key2": "value2"
  },
  "timestamp": "2024-01-10T15:04:05.000Z",
  "caller": "/path/to/file.go:42"
}
```

| フィールド | 説明 |
|-----------|------|
| `level` | ログレベル（DEBUG, INFO, WARN, ERROR, FATAL, PANIC） |
| `msg` | ログメッセージ（文字列またはJSONオブジェクト） |
| `meta` | SetPrefix で設定したメタデータ |
| `timestamp` | ISO 8601 形式のタイムスタンプ |
| `caller` | 呼び出し元のファイル名と行番号 |

## 依存関係

- [github.com/rs/zerolog](https://github.com/rs/zerolog) - 構造化ロガー
- [golang.org/x/xerrors](https://pkg.go.dev/golang.org/x/xerrors) - エラーハンドリング（テスト用）

## ライセンス

MIT License
