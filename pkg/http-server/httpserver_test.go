package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpServer(t *testing.T) {
	port := 8090
	t.Run("should http server include correct values", func(t *testing.T) {

		routes := make(map[string]http.HandlerFunc)
		httpServer := NewHttpServer(port, routes)
		assert.Equal(t, httpServer.Port, port)
		assert.NotNil(t, httpServer.Routes)
	})
	t.Run("should call homepage and get correct response and headers", func(t *testing.T) {

		routes := make(map[string]http.HandlerFunc)
		routes["/"] = HandleJson(func(r *http.Request) (map[string]interface{}, *HttpError) {
			return map[string]interface{}{"msg": "hello"}, nil
		})
		httpServer := NewHttpServer(port, routes)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		assert.NotNil(t, httpServer.Handler)
		httpServer.Handler.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code, "expected status 200")

		ct := rec.Header().Get("Content-Type")
		assert.Equal(t, "application/json", ct, "expected Content-Type application/json")

		var body map[string]interface{}
		err := json.NewDecoder(rec.Body).Decode(&body)
		assert.NoError(t, err, "failed to decode response body")

		msg, ok := body["msg"]
		assert.True(t, ok, "response missing 'msg' field")
		assert.Equal(t, "hello", msg, "expected msg to be 'hello'")
	})

	t.Run("should call homepage and get error with HttpError statusCode", func(t *testing.T) {

		routes := make(map[string]http.HandlerFunc)
		routes["/"] = HandleJson(func(r *http.Request) (map[string]interface{}, *HttpError) {
			return nil, &HttpError{StatusCode: 500, Message: "failed"}
		})
		httpServer := NewHttpServer(port, routes)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		assert.NotNil(t, httpServer.Handler)
		httpServer.Handler.ServeHTTP(rec, req)
		assert.Equal(t, 500, rec.Code, "expected status 500")

		ct := rec.Header().Get("Content-Type")
		assert.Equal(t, "application/json", ct, "expected Content-Type application/json")

		var body map[string]interface{}
		err := json.NewDecoder(rec.Body).Decode(&body)
		assert.NoError(t, err, "failed to decode response body")

		msg, ok := body["error"]
		assert.True(t, ok, "response missing 'error' field")
		assert.Equal(t, "failed", msg, "expected msg to be 'failed'")
	})
	t.Run("should call wrong path and get 404", func(t *testing.T) {
		routes := make(map[string]http.HandlerFunc)
		httpServer := NewHttpServer(port, routes)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		assert.NotNil(t, httpServer.Handler)
		httpServer.Handler.ServeHTTP(rec, req)
		assert.Equal(t, 404, rec.Code, "expected status 404")
	})
}
