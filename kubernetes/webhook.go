package main

import (
    "strconv"
    "net/http"
    "log"
    "bytes"
    "encoding/json"

    "github.com/julienschmidt/httprouter"
)

type CallbackBody struct{}

type Webhook struct {
    CallbackURL string      `json="callback_uri"`
    PushData    PushData    `json="push_data"`
    Repositary  Repositary  `json="repositary"`
}

type PushData struct {
    Images      []string    `json="images"`
    Pusher      string      `json="pusher"`
    Tag         string      `json="tag"`
}

type Repositary struct {
    Name        string      `json="name"`
    Namespace   string      `json="namespace"`
    Owner       string      `json="owner"`
    RepoName    string      `json="repo_name"`
    RepoURL     string      `json="repo_url"`
}

func (s *Server) handleWebhook() httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var webhook Webhook
        if err := decodeBody(r, &webhook); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read webhook from the request")
            return
        }
        webhook.log()
        cb, err := json.Marshal(&CallbackBody{})
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, "failed to create validation callback body")
            return
        }
        resp, err := http.Post(webhook.CallbackURL, "application/json", bytes.NewBuffer(cb))
        if err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to validate webhook origin")
            return
        }
        defer resp.Body.Close()
        respond(w, r, http.StatusOK, nil)
        log.Println("Successfully validate the request")
    }
}

func (w Webhook) log() {
    log.Println("Callback: " + w.CallbackURL)
    for index, image := range w.PushData.Images {
        log.Println("Images[" + strconv.Itoa(index) + "]: " + image)
    }
    log.Println("Pusher: " + w.PushData.Pusher)
    log.Println("Tag: " + w.PushData.Tag)
    log.Println("Name: " + w.Repositary.Name)
    log.Println("Namespace: " + w.Repositary.Namespace)
    log.Println("Owner: " + w.Repositary.Owner)
    log.Println("RepoName: " + w.Repositary.RepoName)
    log.Println("RepoURL: " + w.Repositary.RepoURL)
}
