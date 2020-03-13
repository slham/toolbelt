package l

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"log"
	"net/http"
	"toolbelt/constants"
)

type Level string

type Log struct {
	Data    Data
	Message string
}

type Data struct {
	TranId string
	UserId string
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

func Initialize(lvl string) bool {
	if m, ok := levels[Level(lvl)]; ok {
		mode = m
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
		data := getData(r)
		ctx := r.Context()
		ctx = PushData("reqCtx", data, ctx)
		r = r.WithContext(ctx)
		Debug(r.Context(), fmt.Sprintf("method:%s,url:%s", r.Method, r.URL.Path))
		Debug(r.Context(), fmt.Sprintf("headers:%v", r.Header))
		h.ServeHTTP(w, r)
	})
}

func getData(r *http.Request) Data {
	tranId := uuid.New()
	userId := r.Header.Get(constants.UserId)
	//role := r.Header.Get(constants.Role)
	data := Data{TranId: tranId, UserId: userId}
	return data
}

func PushData(key string, data Data, ctx context.Context) context.Context {
	return context.WithValue(ctx, key, data)
}

func PullData(key string, ctx context.Context) (Data, bool) {
	d, ok := ctx.Value(key).(Data)
	return d, ok
}

func Debug(ctx context.Context, message string, args ...interface{}) {
	if mode >= levels[DEBUG] {
		return
	}
	data, _ := getCtx(ctx)
	l(DEBUG, data.TranId, data.UserId, fmt.Sprintf(message, args...))
}

func Info(ctx context.Context, message string, args ...interface{}) {
	if mode >= levels[INFO] {
		return
	}
	data, _ := getCtx(ctx)
	l(INFO, data.TranId, data.UserId, fmt.Sprintf(message, args...))
}

func Warn(ctx context.Context, message string, args ...interface{}) {
	if mode >= levels[WARN] {
		return
	}
	data, _ := getCtx(ctx)
	l(WARN, data.TranId, data.UserId, fmt.Sprintf(message, args...))
}

func Error(ctx context.Context, message string, args ...interface{}) {
	if mode >= levels[ERROR] {
		return
	}
	data, _ := getCtx(ctx)
	l(ERROR, data.TranId, data.UserId, fmt.Sprintf(message, args...))
}

func l(lvl Level, tranId, userName, message string) {
	l := &Log{Data: Data{UserId: userName, TranId: tranId}, Message: message}
	j, err := json.Marshal(l)
	if err != nil {
		log.Printf("%v::%v", ERROR, err.Error())
	}
	log.Printf("%v::%v", lvl, string(j))
}

func getCtx(ctx context.Context) (Data, bool) {
	return PullData(constants.ReqCtx, ctx)
}
