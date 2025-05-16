package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func HandleJson(method func(r *http.Request) (map[string]interface{}, *HttpError)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		resp, httpErr := method(r)

		if httpErr != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpErr.StatusCode)
			json.NewEncoder(w).Encode(map[string]string{
				"error": httpErr.Message,
			})
			log.Printf("[HTTP] %s %s -> %d (error: %s)", r.Method, r.URL.Path, httpErr.StatusCode, httpErr.Message)
			return
		}

		respJson, err := json.Marshal(resp)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			log.Printf("[HTTP] %s %s -> 500 (marshal error: %s)", r.Method, r.URL.Path, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		end := time.Since(start)
		log.Printf("[HTTP] %s %s -> 200; took %dms", r.Method, r.URL.Path, end.Milliseconds())
	}
}
