package l

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/slham/toolbelt/constants"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	tables := []struct {
		level Level
		succ  bool
	}{
		{"DEBUG", true},
		{"INFO", true},
		{"WARN", true},
		{"ERROR", true},
		{"Master Roshi", false},
	}

	for _, table := range tables {
		result := Initialize(table.level)
		assert.Equal(t, table.succ, result)
	}
}

func TestLogging(t *testing.T) {
	var str bytes.Buffer
	log.SetOutput(&str)
	Initialize("DEBUG")

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("inside dummy handler")
	})

	handlerToTest := Logging(nextHandler)

	req := httptest.NewRequest("GET", "http://test/ing", nil)
	req.Header.Set(constants.UserId, "Willy-Wonka")
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}
