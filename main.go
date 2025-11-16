package main

import (
	"fmt"
	"github.com/frederickrohn/gogogo/user"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r* http.Request){
	u:=user.User{
		ID: 1,
		Name: "Fred",
	}
	greeting:=user.Greet(u)
	fmt.Fprintln(w, greeting)
}



func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}