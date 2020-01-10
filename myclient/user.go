package myclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type UserService struct {
	Client *http.Client
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c *UserService) CreateUser(name string) (*User, error) {
	data, err := json.Marshal(map[string]string{
		"name": name,
	})

	if err != nil {
		return nil, err
	}

	res, err := c.Client.Post("https://api.exmaple.com/users", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var user User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
