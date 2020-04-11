package main

type Image struct {
    Owner   string
    Name    string
    Repo    string
    Tag     string
}

func (webhook *Webhook) imageFromWebhook() *Image {
    return &Image{
        webhook.Repository.Owner,
        webhook.Repository.Name,
        webhook.Repository.RepoName,
        webhook.PushData.Tag,
    }
}
