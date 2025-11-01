package models

import "fmt"

type User struct {
	FirstName string
	Role string
}

func (u User) HelloUser() {
	fmt.Println("Hello -", u.FirstName)
}