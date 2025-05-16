package main

import (
	"net/http"
	httpserver "signit/pkg/http-server"
)

func main() {
	routes := make(map[string]http.HandlerFunc)
	homeFunc := func(r *http.Request) (map[string]interface{}, *httpserver.HttpError) {
		resp := map[string]interface{}{
			"message": "hello world",
		}
		return resp, nil
	}
	routes["/"] = httpserver.HandleHttp(homeFunc)
	server := httpserver.NewHttpServer(8080, routes)
	server.Init()
}
