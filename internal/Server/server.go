package Server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	http.Server
}

func NewServer(m *mux.Router) *Server {
	s := Server{}
	s.Addr = ":8080"
	s.Handler = m
	return &s
}
