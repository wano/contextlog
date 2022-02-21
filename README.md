# contextlog

[github.com/labstack/gommon/log](https://github.com/labstack/gommon/tree/master/log) の interface を用いた [zerolog](https://github.com/rs/zerolog) のラッパー。

## なぜ実装がzelologなのか?
* github.com/labstack/gommon/log　は カスタマイズ性に乏しい
* 開発が止まっている
* スレッドロック等、zap以降のロガーをつかっていた方が[良さげ](https://zenn.dev/moriyoshi/articles/1af0659e29d727)


## install

```go
go get -u github.com/wano/contextlog/...
```

## Usage

global  

```go

clog.Debug(`@@@@`)
// {"level":"DEBUG","caller":"/Users/xxxx/go/src/github.com/wano/contextlog/clog/test/glog_test.go:11","message":"@@@@"}

clog.SetGlobalLevel(zerolog.ErrorLevel)
clog.Info(`@@@@F`)
// no output
```


with Context  

```go

clog.SetGlobalLevel(zerolog.InfoLevel)
ctx := context.Background()
logger := clog.NewContextLogger()
logger.SetPrefix(`app_name` , `app_server`)
logger.SetPrefix(`context_id` , `random_value`)
ctx = logger.WithContext(ctx)

clog.Ctx(ctx).Infof(`message %s` , `!!`)

// {"level":"INFO","app_name":"app_server","context_id":"random_value","caller":"/Users/xxxx/go/src/github.com/wano/contextlog/clog/test/glog_test.go:26","message":"message !!"}

```

## other  features
[interface](./clog/iface.go)