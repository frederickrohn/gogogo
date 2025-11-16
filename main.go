package main

import (
	"fmt"
	"github.com/frederickrohn/gogogo/user"
	"net/http"
	"encoding/json"
)

var users = []user.User{
    {ID: 1, Name: "Fred"},
    {ID: 2, Name: "Alice"},
}

func helloHandler(w http.ResponseWriter, r* http.Request){
	u:=user.User{
		ID: 1,
		Name: "Fred",
	}
	greeting:=user.Greet(u)
	fmt.Fprintln(w, greeting)
}

func usersHandler(w http.ResponseWriter, r* http.Request){

	switch r.Method{
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
		case http.MethodPost:
			http.Error(w, "POST /users is not implemented", http.StatusNotImplemented)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	
}



func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", usersHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}