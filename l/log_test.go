package l

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/slham/toolbelt/constants"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInitialize(t *testing.T) {
	tables := []struct {
		level Level
		succ  bool
		m     int8
	}{
		{"DEBUG", true, int8(0)},
		{"INFO", true, int8(1)},
		{"WARN", true, int8(2)},
		{"ERROR", true, int8(3)},
		{"Master Roshi", false, int8(3)},
	}

	for _, table := range tables {
		result := Initialize(table.level)
		assert.Equal(t, table.succ, result)
		assert.Equal(t, table.m, mode)
	}
}

func TestLogging(t *testing.T) {
	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(0)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("inside dummy handler")
	})

	handlerToTest := Logging(nextHandler)

	req := httptest.NewRequest("GET", "http://test/ing", nil)
	req.Header.Set(constants.UserId, "Willy-Wonka")
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)

	logs := strings.Split(str.String(), "\n")
	assert.Equal(t, 3, len(logs))

	first := cleanLogOutput(logs[0])
	second := cleanLogOutput(logs[1])
	third := cleanLogOutput(logs[2])
	assert.Equal(t, "", third)

	var firstLog Log
	var secondLog Log
	err := json.Unmarshal([]byte(first), &firstLog)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal([]byte(second), &secondLog)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, DEBUG, firstLog.Level)
	assert.Equal(t, "method:GET,url:/ing", firstLog.Message)
	assert.NotEqual(t, "", firstLog.TranId)
	assert.Equal(t, DEBUG, secondLog.Level)
	assert.Equal(t, "headers:map[Userid:[Willy-Wonka]]", secondLog.Message)
	assert.NotEqual(t, "", secondLog.TranId)
}

func TestGetCtx(t *testing.T) {
	ctx := dummyContext()
	tranId := getCtx(ctx)

	assert.Equal(t, "tranId", tranId)
}

func TestDebug(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(0)
	Debug(ctx, "dummy message form %s to be formatted", "Teemo")

	assert.Equal(t, "{\"level\":\"DEBUG\",\"tranId\":\"tranId\",\"message\":\"dummy message form Teemo to be formatted\"}", cleanLogOutput(str.String()))
}

func TestDebugNoLog(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(1)
	Debug(ctx, "dummy message form %s to be formatted", "Teemo")

	assert.Equal(t, "", cleanLogOutput(str.String()))
}

func TestDebugNoCtx(t *testing.T) {
	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(0)
	Debug(nil, "dummy message form %s to be formatted", "Teemo")

	assert.Equal(t, "{\"level\":\"DEBUG\",\"tranId\":\"\",\"message\":\"dummy message form Teemo to be formatted\"}", cleanLogOutput(str.String()))
}

func TestInfo(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(1)
	Info(ctx, "dummy message form %s to be formatted", "Krillin")

	assert.Equal(t, "{\"level\":\"INFO\",\"tranId\":\"tranId\",\"message\":\"dummy message form Krillin to be formatted\"}", cleanLogOutput(str.String()))
}

func TestInfoNoLog(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(2)
	Info(ctx, "dummy message form %s to be formatted", "Krillin")

	assert.Equal(t, "", cleanLogOutput(str.String()))
}

func TestInfoNoCtx(t *testing.T) {
	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(1)
	Info(nil, "dummy message form %s to be formatted", "Krillin")

	assert.Equal(t, "{\"level\":\"INFO\",\"tranId\":\"\",\"message\":\"dummy message form Krillin to be formatted\"}", cleanLogOutput(str.String()))
}

func TestWarn(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(2)
	Warn(ctx, "dummy message form %s to be formatted", "Mickey Mouse")

	assert.Equal(t, "{\"level\":\"WARN\",\"tranId\":\"tranId\",\"message\":\"dummy message form Mickey Mouse to be formatted\"}", cleanLogOutput(str.String()))
}

func TestWarnNoLog(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(3)
	Warn(ctx, "dummy message form %s to be formatted", "Mickey Mouse")

	assert.Equal(t, "", cleanLogOutput(str.String()))
}

func TestWarnNoCtx(t *testing.T) {
	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(2)
	Warn(nil, "dummy message form %s to be formatted", "Mickey Mouse")

	assert.Equal(t, "{\"level\":\"WARN\",\"tranId\":\"\",\"message\":\"dummy message form Mickey Mouse to be formatted\"}", cleanLogOutput(str.String()))
}

func TestError(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(3)
	Error(ctx, "dummy message form %s to be formatted", "Captain Planet")

	assert.Equal(t, "{\"level\":\"ERROR\",\"tranId\":\"tranId\",\"message\":\"dummy message form Captain Planet to be formatted\"}", cleanLogOutput(str.String()))
}

func TestErrorNoLog(t *testing.T) {
	ctx := dummyContext()

	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(4)
	Error(ctx, "dummy message form %s to be formatted", "Captain Planet")

	assert.Equal(t, "", cleanLogOutput(str.String()))
}

func TestErrorNoCtx(t *testing.T) {
	var str bytes.Buffer
	log.SetOutput(&str)
	mode = int8(3)
	Error(nil, "dummy message form %s to be formatted", "Captain Planet")

	assert.Equal(t, "{\"level\":\"ERROR\",\"tranId\":\"\",\"message\":\"dummy message form Captain Planet to be formatted\"}", cleanLogOutput(str.String()))
}

func cleanLogOutput(s string) string {
	dateRemoved := constants.YyyymmddHhmmss.ReplaceAllString(s, "")
	return strings.Trim(dateRemoved, " \n")
}

func dummyContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.ReqCtx, "tranId")
	return ctx
}
