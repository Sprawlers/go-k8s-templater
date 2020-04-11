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
    CallbackURL string      `json:"callback_url"`
    PushData    PushData    `json:"push_data"`
    Repository  Repository  `json:"repository"`
}

type PushData struct {
    Images      []string    `json:"images"`
    Pusher      string      `json:"pusher"`
    Tag         string      `json:"tag"`
}

type Repository struct {
    Name        string      `json:"name"`
    Namespace   string      `json:"namespace"`
    Owner       string      `json:"owner"`
    RepoName    string      `json:"repo_name"`
    RepoURL     string      `json:"repo_url"`
}

func (s *Server) handleWebhook() httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var webhook Webhook
        if err := json.NewDecoder(r.Body).Decode(&webhook); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read webhook from the request")
            return
        }
        defer r.Body.Close()
        webhook.logs()
        newImage := webhook.imageFromWebhook()
        if len(newImage.Tag) < 40 {
            return 
        }
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
        result, err := s.updatePod("dev", webhook.imageFromWebhook())
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, "failed to update the coressponding deployment")
            return
        }
        log.Println("Updated deployment %q.", result.GetObjectMeta().GetName())
        respond(w, r, http.StatusOK, nil)
    }
}

func (w Webhook) logs() {
    log.Println("-------------------------")
    for index, image := range w.PushData.Images {
        log.Println("Images[" + strconv.Itoa(index) + "]: " + image)
    }
    log.Println("Pusher: " + w.PushData.Pusher)
    log.Println("Tag: " + w.PushData.Tag)
    log.Println("Name: " + w.Repository.Name)
    log.Println("Namespace: " + w.Repository.Namespace)
    log.Println("Owner: " + w.Repository.Owner)
    log.Println("RepoName: " + w.Repository.RepoName)
    log.Println("RepoURL: " + w.Repository.RepoURL)
}
