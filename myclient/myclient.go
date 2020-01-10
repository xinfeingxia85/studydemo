package myclient

import (
	"net/http"
)

type Client struct {
	Client *http.Client

	User *UserService
}

func NewClient() *Client {
	c := &Client{
		Client: &http.Client{},
	}

	c.User = &UserService{c.Client}

	return c
}
