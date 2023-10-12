package log

import (
	"context"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var once sync.Once

type Options struct {
	DisableCaller bool
	Level         string
	TimeFormat    string
}

func Init(o *Options) {
	once.Do(func() {
		var w zerolog.LevelWriter
		w = zerolog.MultiLevelWriter(os.Stdout)

		if !o.DisableCaller {
			log.Logger = zerolog.New(w).With().Timestamp().Caller().Logger()
		}

		zerolog.TimeFieldFormat = o.TimeFormat

		switch o.Level {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	})
}

type ctxLogKeyType string

var ctxLogKey ctxLogKeyType = "logFields"

// WithLogValues 对传入的ctx加入items确定的key value
func WithLogValues(ctx context.Context, items ...string) context.Context {
	if len(items) == 0 {
		return ctx
	}

	logFields := map[string]string{}
	for k, v := range fromCtxLogItems(ctx) {
		logFields[k] = v
	}
	for i := 0; i+1 < len(items); i += 2 {
		logFields[items[i]] = items[i+1]
	}

	return context.WithValue(ctx, ctxLogKey, logFields)
}

func fromCtxLogItems(ctx context.Context) map[string]string {
	if ctx == nil {
		return map[string]string{}
	}
	raw := ctx.Value(ctxLogKey)
	if raw == nil {
		return map[string]string{}
	}
	return raw.(map[string]string)
}

func WithLogContext(ctx context.Context, e *zerolog.Event) *zerolog.Event {
	if ctx == context.Background() {
		return e
	}

	logFields := fromCtxLogItems(ctx)
	if len(logFields) == 0 {
		return e
	}

	for k, v := range logFields {
		e = e.Str(k, v)
	}

	return e
}

var nilCtx = context.Background()

func Debug(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Debug())
}

func DebugWithoutCtx() *zerolog.Event {
	return WithLogContext(nilCtx, log.Debug())
}

func Info(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Info())
}

func InfoWithoutCtx() *zerolog.Event {
	return WithLogContext(nilCtx, log.Info())
}

func Warn(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Warn())
}

func WarnWithoutCtx() *zerolog.Event {
	return WithLogContext(nilCtx, log.Warn())
}

func WarnErr(ctx context.Context, err error) *zerolog.Event {
	return Warn(ctx).Err(err)
}

func WarnErrWithoutCtx(err error) *zerolog.Event {
	return Warn(nilCtx).Err(err)
}

func Error(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Error())
}

func ErrorWithoutCtx() *zerolog.Event {
	return WithLogContext(nilCtx, log.Error())
}

func Err(ctx context.Context, err error) *zerolog.Event {
	return WithLogContext(ctx, log.Err(err))
}

func ErrWithoutCtx(err error) *zerolog.Event {
	return WithLogContext(nilCtx, log.Err(err))
}

func Fatal(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Fatal())
}

func FatalWithoutCtx(ctx context.Context) *zerolog.Event {
	return WithLogContext(nilCtx, log.Fatal())
}
