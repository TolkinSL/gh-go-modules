package main

import "fmt"
import "github.com/TolkinSL/gh-go-modules/pkg/models"

func main() {
    fmt.Println("Hello from gh-go-modules!")

		user1 := models.User{
			FirstName: "Vasia",
			Role: "admin",
		}

		fmt.Printf("%+v\n", user1)
}