package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func setupGin(recorder *httptest.ResponseRecorder) *gin.Context {
	context, _ := gin.CreateTestContext(recorder)
	return context
}

func TestHealthHandler_ping(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	w := httptest.NewRecorder()
	context := setupGin(w)
	argsData := args{
		ctx: context,
	}

	tests := []struct {
		name   string
		args   args
		expect interface{}
	}{
		{name: "Should return pong", args: argsData, expect: gin.H{"data": "pong"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HealthHandler{}
			h.ping(tt.args.ctx)
			assert.Equal(t, 200, w.Code)
			marshal, _ := json.Marshal(tt.expect)
			assert.Equal(t, string(marshal), w.Body.String())
		})
	}
}
