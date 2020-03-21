package l

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"github.com/slham/toolbelt/constants"
	"log"
	"net/http"
)

type Level string

type Log struct {
	Level   Level  `json:"level"`
	TranId  string `json:"tranId"`
	Message string `json:"message"`
}

const (
	DEBUG Level = "DEBUG"
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
)

var levels = map[Level]int8{
	DEBUG: int8(0),
	INFO:  int8(1),
	WARN:  int8(2),
	ERROR: int8(3),
}

var mode int8

func Initialize(lvl Level) bool {
	if m, ok := levels[lvl]; ok {
		mode = m
		log.Println(fmt.Sprintf("logging level set to %v", lvl))
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
		ctx = context.WithValue(ctx, constants.ReqCtx, uuid.New())
		r = r.WithContext(ctx)
		Debug(r.Context(), fmt.Sprintf("method:%s,url:%s", r.Method, r.URL.Path))
		Debug(r.Context(), fmt.Sprintf("headers:%v", r.Header))
		h.ServeHTTP(w, r)
	})
}

func Debug(ctx context.Context, message string, args ...interface{}) {
	if mode > levels[DEBUG] {
		return
	}
	var tranId string

	if ctx != nil {
		tranId = getCtx(ctx)
	}

	l(DEBUG, tranId, fmt.Sprintf(message, args...))
}

func Info(ctx context.Context, message string, args ...interface{}) {
	if mode > levels[INFO] {
		return
	}
	var tranId string

	if ctx != nil {
		tranId = getCtx(ctx)
	}

	l(INFO, tranId, fmt.Sprintf(message, args...))
}

func Warn(ctx context.Context, message string, args ...interface{}) {
	if mode > levels[WARN] {
		return
	}
	var tranId string

	if ctx != nil {
		tranId = getCtx(ctx)
	}

	l(WARN, tranId, fmt.Sprintf(message, args...))
}

func Error(ctx context.Context, message string, args ...interface{}) {
	if mode > levels[ERROR] {
		return
	}
	var tranId string

	if ctx != nil {
		tranId = getCtx(ctx)
	}

	l(ERROR, tranId, fmt.Sprintf(message, args...))
}

func l(lvl Level, tranId, message string) {
	l := &Log{Level: lvl, TranId: tranId, Message: message}
	j, err := json.Marshal(l)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(j))
}

func getCtx(ctx context.Context) string {
	return ctx.Value(constants.ReqCtx).(string)
}
