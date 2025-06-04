package entities

type User struct {
	ID 	 		int    `json:"id"`
	Name   		string `json:"name"`
	LastName 	string `json:"last_name"`
	Email  		string `json:"email"`
	Password 	string `json:"password"`
}