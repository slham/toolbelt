package l

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"log/slog"

	"github.com/pborman/uuid"
	"github.com/slham/toolbelt/constants"
)

type Level string
type ctxKey string

const (
	DEBUG      Level  = "DEBUG"
	INFO       Level  = "INFO"
	WARN       Level  = "WARN"
	ERROR      Level  = "ERROR"
	slogFields ctxKey = "slog_fields"
)

var levels = map[Level]slog.Level{
	DEBUG: slog.LevelDebug,
	INFO:  slog.LevelInfo,
	WARN:  slog.LevelWarn,
	ERROR: slog.LevelError,
}

type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying
// handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	return h.Handler.Handle(ctx, r)
}

// AppendCtx adds an slog attribute to the provided context so that it will be
// included in any Record created with such context
func AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	if v, ok := parent.Value(slogFields).([]slog.Attr); ok {
		v = append(v, attr)
		return context.WithValue(parent, slogFields, v)
	}

	v := []slog.Attr{}
	v = append(v, attr)
	return context.WithValue(parent, slogFields, v)
}

func Initialize(lvl Level) bool {
	if level, ok := levels[lvl]; ok {
		log.Println(fmt.Sprintf("setting logging level to %s", lvl))

		var programLevel = new(slog.LevelVar)
		programLevel.Set(level)

		h := &ContextHandler{slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})}
		logger := slog.New(h)
		slog.SetDefault(logger)
		log.Println(fmt.Sprintf("logging level set to %s", lvl))
		return ok
	} else {
		log.Println(fmt.Sprintf("invalid logging level %s", lvl))
		return ok
	}
}

//Initializes transaction logging
func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = AppendCtx(ctx, slog.String(constants.ReqCtx, uuid.New()))
		r = r.WithContext(ctx)
		slog.DebugContext(ctx, fmt.Sprintf("method:%s,url:%s", r.Method, r.URL.Path))
		slog.DebugContext(ctx, fmt.Sprintf("headers:%v", r.Header))
		h.ServeHTTP(w, r)
	})
}
