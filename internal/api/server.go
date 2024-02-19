package api

import (
	"fmt"
	"net/http"

	"github.com/VoC925/go-testify/internal/api/config"
	"github.com/VoC925/go-testify/internal/handlers"
	"github.com/go-chi/chi/v5"
)

// Структура сервера
type Server struct {
	Host         string
	Port         string
	Router       *chi.Mux
	HandlerRoute handlers.Handler
}

// New возвращет экземпляр структуры сервера
func New(cfg *config.Config, hand handlers.Handler) *Server {
	return &Server{
		Host:         cfg.Server.Host,
		Port:         cfg.Server.Port,
		Router:       chi.NewRouter(),
		HandlerRoute: hand,
	}
}

// Start метод, запускающий сервер
func (s *Server) Start() error {
	s.HandlerRoute.Register(s.Router)
	fmt.Printf("Server start: host -> %s; port -> :%s", s.Host, s.Port)
	connStr := fmt.Sprintf(`%s:%s`, s.Host, s.Port)
	return http.ListenAndServe(connStr, s.Router)
}
