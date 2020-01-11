package model

type User struct {
	Name     string `json:"name"`
	Password string `json:"password,,omitempty"`
}

type Resp struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
