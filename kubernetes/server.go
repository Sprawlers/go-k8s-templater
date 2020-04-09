package main

import (
    "net/http"

    "github.com/rs/cors"
    "github.com/julienschmidt/httprouter"
    "k8s.io/client-go/kubernetes"
)

type Server struct {
    clientset   *kubernetes.Clientset
    router      *httprouter.Router
    cors        *cors.Cors
}

func newServer(clientset *kubernetes.Clientset) *Server {
    s := &Server{clientset, httprouter.New(), cors.New(options)}
    s.routes()
    options := cors.Options{
        AllowedMethods: []string{"GET", "POST"},
    }
    return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.cors.ServeHTTP(w, r, s.router.ServeHTTP)
}
