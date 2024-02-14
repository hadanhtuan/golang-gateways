package model

type LoginSchema struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
