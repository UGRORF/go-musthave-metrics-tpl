package handler

import (
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handler *MetricHandler) *Server {
	r := routerInit(handler)
	httpServer := &http.Server{
		Addr:    port,
		Handler: r,
	}
	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	fmt.Printf("Starting server on port %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
