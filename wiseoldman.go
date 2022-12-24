package wiseoldman

import "net/http"

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		baseURL:    "https://api.wiseoldman.net/v2",
	}
}

type Error struct {
	Message string `json:"message"`
}
