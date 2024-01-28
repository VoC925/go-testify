package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Структура сервера
type Server struct {
	Host   string
	Port   string
	Router *chi.Mux
}

// New возвращет экземпляр структуры сервера
func New(host string, port string) *Server {
	return &Server{
		Host:   host,
		Port:   port,
		Router: chi.NewRouter(),
	}
}

// Start метод, запускающий сервер
func (s *Server) Start() error {
	s.Router.Get("/cafe", mainHandle)
	fmt.Printf("Server start: host -> %s; port -> %s", s.Host, s.Port)
	return http.ListenAndServe(s.Host+s.Port, s.Router)
}
