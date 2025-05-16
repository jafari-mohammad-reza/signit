package main

import (
	"net/http"
	httpserver "signit/pkg/http-server"
)

func main() {
	server := httpserver.NewHttpServer(8080, getRoutes())
	server.Init()
}

func getRoutes() map[string]http.HandlerFunc {
	routes := make(map[string]http.HandlerFunc)
	service := NewService()
	homeFunc := func(r *http.Request) (map[string]interface{}, *httpserver.HttpError) {
		resp := map[string]interface{}{
			"message": "hello world",
		}
		return resp, nil
	}
	routes["/"] = httpserver.HandleJson(homeFunc)
	routes["/generate-keys"] = httpserver.HandleJson(func(r *http.Request) (map[string]interface{}, *httpserver.HttpError) {
		resp := make(map[string]interface{})
		keyPair, err := service.GenerateKeyPair()
		if err != nil {
			return nil, &httpserver.HttpError{StatusCode: 500, Message: err.Error()}
		}
		resp["privateKey"] = keyPair.PrivateKey
		resp["publicKey"] = keyPair.PublicKey
		return resp, nil
	})
	return routes
}
