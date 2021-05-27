package domain

import "fmt"

type User struct {
	Id uint64					`json:"id"`
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	Email string			`json:"email"`
}

func UserDomainHi() {
	fmt.Println("user domain")
}