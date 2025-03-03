package models

type User struct {
	ID       string `json:"id"`
	Nom      string `json:"nom"` 
	Prenom   string	`json:"prenom"`
	Age      int	`json:"age"`
	Contact  string	`json:"contact"`
	Password string	`json:"password"`
}

type Response struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    *User 			`json:"data,omitempty"`
}