package main

import (
    "net/http"
    "context"

    "k8s.io/apimachinery/pkg/api/errors"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) handleHealtCheck() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        respond(w, r, http.StatusOK, nil)
    }
}

func (s *Server) handleTest() http.HandlerFunc {
    return func(w http.ResonseWriter, r *http.Request) {
        pods, err := clientset.CoreV1().Pods("production").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, err)
            return
        }


    }
}
