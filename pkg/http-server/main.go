package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Route map[string]http.HandlerFunc
type HttpServer struct {
	Port   int
	Routes Route
}

func NewHttpServer(port int, routes Route) *HttpServer {
	return &HttpServer{
		Port:   port,
		Routes: routes,
	}
}

func (h *HttpServer) Init() error {
	mux := http.NewServeMux()
	for route, handler := range h.Routes {
	mux.HandleFunc(route, handler)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", h.Port),
		Handler:      mux,
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
