package models

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Section string `json:"section"`
	Role    string `json:"roles"`
}
