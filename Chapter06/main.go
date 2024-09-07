package main

import "fmt"

func s_main() {
	mensur := User{"Mensur", "mensur@gmail.com", true, 15}
	mensur.GetStatus()
	mensur.SetEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) SetEmail() {
	u.Email += "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
