package models

type Login struct {
	Username string `json:"username"`
	Pin      string `json:"pin"`
}
