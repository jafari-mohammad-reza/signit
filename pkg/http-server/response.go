package httpserver

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func HandleHttp(method func(r *http.Request) (map[string]interface{}, *HttpError)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, httpErr := method(r)
		if httpErr != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpErr.StatusCode)
			json.NewEncoder(w).Encode(map[string]string{
				"error": httpErr.Message,
			})
			return
		}

		respJson, err := json.Marshal(resp)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	}
}
