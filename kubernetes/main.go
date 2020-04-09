package main

import (
    "log"
    "net/http"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

func main() {
    config, err := rest.InClusterConfig()
    if err != nil {
        log.Fatalln(err)
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalln(err)
    }
    s := newServer(clientset)
    http.ListenAndServe(":80", s)
}
