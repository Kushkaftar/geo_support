package GEO_Support

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

// запуск сервера
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

// выход из приложения
func (s *Server) namShutdown(ctx context.Context) error {
	return  s.httpServer.Shutdown(ctx)
}
