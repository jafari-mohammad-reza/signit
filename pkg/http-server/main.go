package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Route map[string]http.HandlerFunc
type HttpServer struct {
	Port    int
	Routes  Route
	Handler http.Handler
}

func NewHttpServer(port int, routes Route) *HttpServer {
	mux := http.NewServeMux()
	for route, handler := range routes {
		mux.HandleFunc(route, handler)
	}
	return &HttpServer{
		Port:    port,
		Routes:  routes,
		Handler: mux,
	}
}

func (h *HttpServer) Init() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", h.Port),
		Handler:      h.Handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server starting on :%d\n", h.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	select {}
}
