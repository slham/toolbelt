package l

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"toolbelt/constants"
)

func TestGetData(t *testing.T) {
	userId := "userId"

	r := http.Request{Header: http.Header{}}
	r.Header.Set(constants.UserId, userId)

	data := getData(&r)
	assert.Equal(t, data.UserId, userId)
	assert.False(t, data.TranId == "", "tranId not set")
}

func TestPushData(t *testing.T) {
	d := Data{TranId: "tranId",  UserId: "subId"}
	key := "testKey"
	ctx := context.Background()

	ctx = PushData(key, d, ctx)

	assert.Equal(t, ctx.Value(key).(Data), d)
}

func TestPullData(t *testing.T) {
	d := Data{TranId: "tranId",  UserId: "subId"}
	key := "testKey"
	ctx := context.Background()

	ctx = context.WithValue(ctx, key, d)

	data, ok := PullData(key, ctx)

	assert.True(t, ok, "error pulling data")
	assert.Equal(t, d, data)
}

func TestInitialize(t *testing.T) {
	tables := []struct{
		level string
		succ bool
		m int8
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

func TestGetCtx(t *testing.T) {
	d := Data{TranId: "tranId",  UserId: "subId"}
	ctx := context.Background()

	ctx = context.WithValue(ctx, constants.ReqCtx, d)

	data, ok := getCtx(ctx)

	assert.True(t, ok, "error getting context")
	assert.Equal(t, d, data)
}
