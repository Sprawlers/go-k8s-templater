package main

import (
    "net/http"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) handleHealtCheck() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        respond(w, r, http.StatusOK, nil)
    }
}

func (s *Server) handleTest() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        _, err := s.clientset.CoreV1().Pods("production").List(metav1.ListOptions{})
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, err)
            return
        }
    }
}
