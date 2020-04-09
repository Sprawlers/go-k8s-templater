package main

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
)

func (s *Server) routes() {
    s.router.HandlerFunc("GET", "/health", s.handleHealtCheck())
    s.router.HadnlerFunc("GET", "/test", s.handleTest())
}
