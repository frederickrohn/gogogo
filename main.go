package main

import "fmt"

type User struct {
	ID int
	Name string
}

func greet(u User) string{
	str := "Hi! I'm " + u.Name
	return str
}

func main() {

	user:= User{
		ID: 1,
		Name: "Fred",
	}
	greeting:= greet(user)

	fmt.Println(greeting)
}