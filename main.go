package main

import (
	"fmt"
	"github.com/frederickrohn/gogogo/user"
)

func main() {

	u:= user.User{
		ID: 1,
		Name: "Fred",
	}
	greeting:= user.Greet(u)

	fmt.Println(greeting)
}